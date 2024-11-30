package discovery

import (
	"context"
	"fmt"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func clientUnaryInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	md.Append("authorization", "Bearer hardcoded")

	newCtx := metadata.NewOutgoingContext(ctx, md)

	return invoker(newCtx, method, req, reply, cc, opts...)
}

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addressList, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	if len(addressList) == 0 {
		return nil, fmt.Errorf("no available instances found for service: %s", serviceName)
	}

	address := addressList[rand.Intn(len(addressList))]

	// TODO: grpc.NewClient was taking around 10s to process a request,
	// grpc.Dial is working as expected, so for now using grpc.Dial
	// Will revisit it in future
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientUnaryInterceptor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service %s: %v", serviceName, err)
	}

	return conn, nil
}
