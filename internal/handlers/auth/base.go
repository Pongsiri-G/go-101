package auth

import (
	"context"

	"github.com/graphzc/go-clean-template/internal/dto"
	"github.com/graphzc/go-clean-template/internal/services/user"
)

type Handler interface {
	Register(ctx context.Context, req dto.UserRegisterRequest) (*dto.MessageResponse, error)
}

type handler struct {
	service user.Service
}

// @WireSet("Handler")
func New(service user.Service) Handler {
	return &handler{
		service: service,
	}
}

// Register implements Handler.
func (h *handler) Register(ctx context.Context, req dto.UserRegisterRequest) (*dto.MessageResponse, error) {
	userRegisterInput := &user.UserRegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := h.service.Register(ctx, userRegisterInput); err != nil {
		return nil, err
	}

	return &dto.MessageResponse{
		Message: "User registered successfully",
	}, nil

}
