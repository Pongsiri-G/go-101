package server

import (
	"fmt"

	"github.com/graphzc/go-clean-template/internal/config"
	"github.com/graphzc/go-clean-template/internal/handlers"
	"github.com/graphzc/go-clean-template/internal/router"
	"github.com/graphzc/go-clean-template/internal/utils/servererr"
	"github.com/graphzc/go-clean-template/internal/utils/validator"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	config   *config.Config
	handlers *handlers.Handlers
}

func NewEchoServer(
	config *config.Config,
	handlers *handlers.Handlers,
) *EchoServer {
	return &EchoServer{
		config:   config,
		handlers: handlers,
	}
}

func (s *EchoServer) Start() error {
	e := echo.New()

	e.Validator = validator.NewValidator()

	e.HTTPErrorHandler = servererr.EchoHTTPErrorHandler

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     s.config.CORS.AllowOrigins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	router := router.NewRouter(e, s.handlers)

	router.RegisterAPIRoutes()

	return e.Start(fmt.Sprintf(":%s", s.config.Port))
}
