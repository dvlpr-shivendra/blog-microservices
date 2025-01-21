package main

import (
	"blog-services/common"
	"blog-services/common/discovery"
	"blog-services/common/discovery/consul"
	"blog-services/gateway/gateway"
	"context"
	"log"
	"net/http"
	"time"
)

var (
	serviceName = "gateway"
	httpAddr    = common.Env("HTTP_ADDR", ":8080")
	consulAddr  = common.Env("CONSUL_ADDR", "localhost:8500")
	jaegerAddr  = common.Env("JAEGER_ADDR", "localhost:4318")
)

func main() {
	err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr)

	if err != nil {
		log.Fatal("failed to set global tracer")
	}

	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
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
	mux := http.NewServeMux()

	postsGateway := gateway.NewGRPCGateway(registry)

	handler := NewHandler(postsGateway)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
