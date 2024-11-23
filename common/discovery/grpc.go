package discovery

import (
	"context"
	"fmt"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addressList, err := registry.Discover(ctx, serviceName)

	if err != nil {
		return nil, err
	}

	if len(addressList) == 0 {
		return nil, fmt.Errorf("no available instances found for service: %s", serviceName)
	}

	return grpc.NewClient(addressList[rand.Intn(len(addressList))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
