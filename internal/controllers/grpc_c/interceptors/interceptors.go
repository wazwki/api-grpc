package interceptors

import (
	"context"
	"strings"
	"time"

	"github.com/wazwki/api-grpc/internal/config"
	"github.com/wazwki/api-grpc/pkg/jwtutil"
	"github.com/wazwki/api-grpc/pkg/logger"
	"github.com/wazwki/api-grpc/pkg/metrics"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func JWTInterceptor(cfg *config.Config, jwt *jwtutil.JWTUtil) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == "/api.v1.Name/HealthCheck" {
			return handler(ctx, req)
		}
		/*
			if info.FullMethod == "/api.v1.Name/HealthCheck" {
				token, err := jwt.GenerateAccessToken(ctx)
				if err != nil {
					return nil, status.Error(codes.Unauthenticated, "failed to generate access token")
				}

				ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
					"Authorization": "Bearer " + token,
				}))
			}
		*/
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md.Get("Authorization")
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		_, err := jwt.ValidateToken(ctx, tokenStr)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		return handler(ctx, req)
	}
}

func MetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start).Seconds()

		metrics.ObserveRequestDuration.WithLabelValues(info.FullMethod).Observe(duration)
		return resp, err
	}
}

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info("gRPC request", zap.String("method", info.FullMethod), zap.String("module", "skillsrock"))

		return handler(ctx, req)
	}
}
