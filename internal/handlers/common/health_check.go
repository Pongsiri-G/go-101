package common

import (
	"context"

	"github.com/graphzc/go-clean-template/internal/dto"
)

func (h *handler) HealthCheck(ctx context.Context, _ any) (dto.HealthCheckResponse, error) {
	return dto.HealthCheckResponse{
		Status: "ok",
	}, nil
}
