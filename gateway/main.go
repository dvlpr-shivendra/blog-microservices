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

	"go.uber.org/zap"
)

var (
	serviceName = "gateway"
	httpAddr    = common.Env("HTTP_ADDR", ":8080")
	consulAddr  = common.Env("CONSUL_ADDR", "localhost:8500")
	jaegerAddr  = common.Env("JAEGER_ADDR", "localhost:4318")
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr)

	if err != nil {
		logger.Fatal("could set global tracer", zap.Error(err))
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
		ticker := time.NewTicker(10 * time.Second) // Increase interval
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := registry.HealthCheck(instanceID, serviceName); err != nil {
					log.Println("Health check failed:", err)
				}
			}
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)
	mux := http.NewServeMux()

	postsGateway := gateway.NewGRPCGateway(registry, logger)

	handler := NewHandler(postsGateway, logger)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
