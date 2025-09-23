package router

import (
	"github.com/graphzc/go-clean-template/internal/handlers"
	"github.com/labstack/echo/v4"
)

type Router struct {
	echo     *echo.Echo
	handlers *handlers.Handlers
}

func NewRouter(
	echo *echo.Echo,
	handlers *handlers.Handlers,
) *Router {
	return &Router{
		echo:     echo,
		handlers: handlers,
	}
}
