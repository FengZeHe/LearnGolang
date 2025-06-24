package jwt

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type JWTInterceptor struct {
	jwtToken string
}

func NewJWTInterceptor(jwtToken string) *JWTInterceptor {
	return &JWTInterceptor{jwtToken: jwtToken}
}

func (j *JWTInterceptor) JwtClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := metadata.Pairs("authorization", "Bearer "+j.jwtToken)

		ctx = metadata.NewOutgoingContext(ctx, md)
		// 继续调用链路
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
