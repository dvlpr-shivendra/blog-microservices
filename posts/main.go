package main

import (
	"blog-services/common"
	"blog-services/common/broker"
	"blog-services/common/cache"
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

	"go.uber.org/zap"
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
	jaegerAddr  = common.Env("JAEGER_ADDR", "localhost:4318")
	amqpUser    = common.Env("RABBITMQ_USER", "guest")
	amqpPass    = common.Env("RABBITMQ_PASS", "guest")
	amqpHost    = common.Env("RABBITMQ_HOST", "localhost")
	amqpPort    = common.Env("RABBITMQ_PORT", "5672")
	redisAddr     = common.Env("REDIS_ADDR", "localhost:6379")
	redisPassword = common.Env("REDIS_PASSWORD", "")
	redisDB       = 0 // Default DB Index

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
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	if err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr); err != nil {
		logger.Fatal("could set global tracer", zap.Error(err))
	}

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
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := registry.HealthCheck(instanceID, serviceName); err != nil {
					logger.Warn("health check failed", zap.Error(err))
					time.Sleep(5 * time.Second) // Backoff before retry
				}
			}
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(1000),
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
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
		logger.Fatal("failed to connect to db", zap.Error(err))
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	redisCache := cache.NewRedisCache(redisAddr, redisPassword, redisDB)
	defer redisCache.Close()

	store := NewStore(db)
	svc := NewService(store, redisCache)
	svcWithTelemetry := NewTelemetryMiddleware(svc)
	svcWithLogging := NewLoggingMiddleware(svcWithTelemetry)
	NewGRPCHandler(grpcServer, svcWithLogging)
	consumer := NewConsumer(svcWithLogging)
	consumer.Listen(ch)
	logger.Info("starting HTTP server", zap.String("port", grpcAddr))

	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
