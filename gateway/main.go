package main

import (
	"blog-services/common"
	"blog-services/gateway/gateway"
	"log"
	"net/http"
)

var httpAddr = common.Env("HTTP_ADDR", "localhost:8080")

func main() {
	mux := http.NewServeMux()

	postsGateway := gateway.NewGRPCGateway()

	handler := NewHandler(postsGateway)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
