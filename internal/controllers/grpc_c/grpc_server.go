package grpc_c

import (
	"context"
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
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.JWTInterceptor(cfg, jwt),
			interceptors.MetricsInterceptor(),
			interceptors.LoggerInterceptor(),
		),
	)
	namepb.RegisterNameServiceServer(grpcServer, nameControllers)

	runtimeMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := namepb.RegisterNameServiceHandlerFromEndpoint(context.Background(), runtimeMux, cfg.GRPCPort, opts)
	if err != nil {
		return nil, nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", runtimeMux)

	srv := &http.Server{
		Addr:    cfg.GRPCPort,
		Handler: mux,
	}

	return grpcServer, srv, nil
}
