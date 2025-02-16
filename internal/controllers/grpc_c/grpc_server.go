package grpc_c

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wazwki/api-grpc/api/proto/namepb"
	"github.com/wazwki/api-grpc/internal/config"
	"github.com/wazwki/api-grpc/internal/controllers/grpc_c/interceptors"
	"github.com/wazwki/api-grpc/pkg/jwtutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCServer(cfg *config.Config, nameControllers namepb.NameServiceServer, jwt *jwtutil.JWTUtil) (*grpc.Server, *http.Server, error) {
	var interceptorsUse []grpc.UnaryServerInterceptor

	switch cfg.Debug {
	case true:
		interceptorsUse = []grpc.UnaryServerInterceptor{
			interceptors.LoggerInterceptor(),
			interceptors.MetricsInterceptor(),
		}
	case false:
		interceptorsUse = []grpc.UnaryServerInterceptor{
			interceptors.JWTInterceptor(cfg, jwt),
			interceptors.LoggerInterceptor(),
			interceptors.MetricsInterceptor(),
		}
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptorsUse...))

	namepb.RegisterNameServiceServer(grpcServer, nameControllers)

	runtimeMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := namepb.RegisterNameServiceHandlerFromEndpoint(context.Background(), runtimeMux, fmt.Sprintf("%v:%v", cfg.Host, cfg.GRPCPort), opts)
	if err != nil {
		return nil, nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", runtimeMux)
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Handler: mux,
	}

	return grpcServer, srv, nil
}
