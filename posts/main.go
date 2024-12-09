package main

import (
	"blog-services/common"
	"blog-services/common/discovery"
	"blog-services/common/discovery/consul"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"

	"database/sql"

	_ "github.com/lib/pq"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	serviceName = "posts"
	grpcAddr    = common.Env("GRPC_ADDR", "localhost:2000")
	consulAddr  = common.Env("CONSUL_ADDR", "localhost:8500")
	dbUser      = common.Env("DB_USER", "postgres")
	dbName      = common.Env("DB_NAME", "blog_posts_db")
	dbPassword  = common.Env("DB_PASSWORD", "postgres")
	dbHost      = common.Env("DB_HOST", "localhost")
	dbPort      = common.Env("DB_PORT", "5432")
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Extract metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Println("❌ No metadata found in context")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}

	token, ok := md["authorization"]

	if !ok || len(token) < 1 {
		log.Println("❌ No token found in metadata")
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
	}

	isValid := validateToken(token[0])

	if !isValid {
		log.Printf("❌ Invalid token received: %s", token)
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
	}

	return handler(ctx, req)
}

func validateToken(token string) bool {
	// Add your token validation logic here
	// For example:
	// - Validate JWT token
	// - Check against your auth service
	// - Verify token format and expiration
	return token != ""
}

func main() {
	registry, err := consul.NewRegistry(consulAddr, serviceName)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		panic(err)
	}

	defer l.Close()

	connectionString := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser,
		dbName,
		dbPassword,
		dbHost,
		dbPort,
	)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	svc := NewService(store)

	NewGRPCHandler(grpcServer, svc)

	log.Printf("Starting GRPC server at %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
