package router

import (
	"net/http"

	"github.com/graphzc/go-clean-template/internal/utils/echoutil"
)

func (r *Router) RegisterAPIRoutes() {

	// Health check
	r.echo.GET("/health", echoutil.WrapWithStatus(r.handlers.Common.HealthCheck, http.StatusOK))

	v1Public := r.echo.Group("/api/v1")

	authGroupV1 := v1Public.Group("/auth")
	// ปีกกาช่วยให้อ่านง่ายขึ้นเฉยๆ
	{
		// /api/v1/auth/register
		authGroupV1.POST("/register", echoutil.WrapWithStatus(r.handlers.Auth.Register, http.StatusCreated))
	}
}
