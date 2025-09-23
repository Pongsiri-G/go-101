package router

import (
	"net/http"

	"github.com/graphzc/go-clean-template/internal/utils/echoutil"
)

func (r *Router) RegisterAPIRoutes() {
	// Health check
	r.echo.GET("/health", echoutil.WrapWithStatus(r.handlers.Common.HealthCheck, http.StatusOK))
}
