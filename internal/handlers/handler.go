package handlers

import (
	"github.com/graphzc/go-clean-template/internal/handlers/auth"
	"github.com/graphzc/go-clean-template/internal/handlers/common"
)

type Handlers struct {
	Common common.Handler
	Auth   auth.Handler
}

// @WireSet("Handler")
func NewHandlers(
	commonHandler common.Handler,
	authHandler auth.Handler,

) *Handlers {
	return &Handlers{
		Common: commonHandler,
		Auth:   authHandler,
	}
}
