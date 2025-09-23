package foo

import "github.com/graphzc/go-clean-template/internal/services/foo"

type Handler interface {
}

type handler struct {
	service foo.Service
}

// @WireSet("Handler")
func New(service foo.Service) Handler {
	return &handler{
		service: service,
	}
}
