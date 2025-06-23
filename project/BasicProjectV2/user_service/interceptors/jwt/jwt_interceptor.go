package jwt

import (
	"context"
	"github.com/basicprojectv2/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type JWTInterceptor struct {
	exemptPath map[string]bool
}

func NewJWTInterceptor(exemptPath []string) *JWTInterceptor {
	pathMap := make(map[string]bool)
	for _, path := range exemptPath {
		pathMap[path] = true
	}
	return &JWTInterceptor{
		exemptPath: pathMap,
	}
}

func (j *JWTInterceptor) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if j.exemptPath[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "缺少Metadata")
		}

		// 提取jwt token
		authHeader := md.Get("Authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "缺少Auth")
		}

		//解析jwt token
		parts := strings.Split(authHeader[0], " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return nil, status.Errorf(codes.Unauthenticated, "无效Auth格式")
		}

		tokenString := parts[1]
		if tokenString == "" {
			return nil, status.Errorf(codes.Unauthenticated, "空的JWT令牌")
		}

		clamis, err := jwt.ParseToken(tokenString)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "JWT Token错误")
		}
		ctx = context.WithValue(ctx, "user_id", clamis.UserId)
		return handler(ctx, req)
	}
}
