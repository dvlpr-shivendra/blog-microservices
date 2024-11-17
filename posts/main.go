package main

import (
	"blog-services/common"
	"blog-services/common/discovery"
	"blog-services/common/discovery/consul"
	"context"
	"log"
	"net"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	serviceName = "orders"
	grpcAddr    = common.Env("GRPC_ADDR", "localhost:2000")
	consulAddr  = common.Env("CONSUL_ADDR", "localhost:8500")
)

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

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		panic(err)
	}

	defer l.Close()

	store := NewStore()
	svc := NewService(store)

	NewGRPCHandler(grpcServer, svc)

	log.Printf("Starting GRPC server at %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
