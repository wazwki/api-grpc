package v1

import (
	"context"

	"github.com/wazwki/api-grpc/api/proto/namepb"
	"github.com/wazwki/api-grpc/internal/service"
)

type NameControllers struct {
	namepb.UnimplementedNameServiceServer
	service service.NameServiceInterface
}

func NewNameControllers(service service.NameServiceInterface) namepb.NameServiceServer {
	return &NameControllers{
		service: service,
	}
}

func (s *NameControllers) HealthCheck(ctx context.Context, req *namepb.HealthCheckRequest) (*namepb.HealthCheckResponse, error) {
	return &namepb.HealthCheckResponse{Status: "OK"}, nil
}
