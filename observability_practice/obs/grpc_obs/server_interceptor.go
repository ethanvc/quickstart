package grpc_obs

import (
	"context"
	"google.golang.org/grpc"
)

type ServerInterceptor struct {
}

func CreateServerInterceptor() *ServerInterceptor {
	si := &ServerInterceptor{}
	return si
}

func (si *ServerInterceptor) Process(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	return handler(ctx, req)
}
