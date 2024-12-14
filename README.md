# Blog Microservices

A microservices-based blog platform built with Go, TypeScript, and gRPC.

## Services

### Posts Service (Go)
Handles blog post management including creation, updates, and retrieval.
- Language: Go
- Database: PostgreSQL
- Communication: gRPC

### Comments Service (Go)
Manages comments on blog posts.
- Language: Go
- Database: PostgreSQL
- Communication: gRPC

### Likes Service (TypeScript)
Handles post likes functionality.
- Language: TypeScript/Node.js
- Storage: In-memory (can be extended to use a database)
- Communication: gRPC

### Gateway Service (Go)
API Gateway that routes requests to appropriate microservices.
- Language: Go
- Communication: HTTP REST (external), gRPC (internal)

## Architecture

- Service Discovery: Consul
- Communication Protocol: gRPC
- API Gateway Pattern
- Environment-based Configuration
- Health Checking

## Prerequisites

- Go 1.22+
- Node.js 18+
- Docker and Docker Compose
- PostgreSQL
- Consul