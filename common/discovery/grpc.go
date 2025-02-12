package discovery

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
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

	// Create a resolver that will be used by the round-robin balancer
	resolverScheme := "static"
	resolverBuilder := resolver.Get(resolverScheme)
	if resolverBuilder == nil {
		r := &staticResolver{addresses: addressList}
		resolver.Register(r)
	}

	// Format the target address using the static scheme
	target := fmt.Sprintf("%s:///unused", resolverScheme)

	// Create the client connection with round-robin load balancing
	conn, err := grpc.NewClient(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientUnaryInterceptor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service %s: %v", serviceName, err)
	}

	return conn, nil
}

// staticResolver implements the resolver.Builder interface
type staticResolver struct {
	addresses []string
}

func (r *staticResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	addrs := make([]resolver.Address, len(r.addresses))
	for i, addr := range r.addresses {
		addrs[i] = resolver.Address{Addr: addr}
	}
	cc.UpdateState(resolver.State{Addresses: addrs})
	return &nopResolver{cc: cc}, nil
}

func (r *staticResolver) Scheme() string {
	return "static"
}

// nopResolver implements the resolver.Resolver interface
type nopResolver struct {
	cc resolver.ClientConn
}

func (r *nopResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (r *nopResolver) Close()                                {}
