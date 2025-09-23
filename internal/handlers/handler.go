package handlers

import (
	"github.com/graphzc/go-clean-template/internal/handlers/common"
	"github.com/graphzc/go-clean-template/internal/handlers/foo"
)

type Handlers struct {
	Common common.Handler
	Foo    foo.Handler
}

// @WireSet("Handler")
func NewHandlers(
	commonHandler common.Handler,
	fooHandler foo.Handler,
) *Handlers {
	return &Handlers{
		Common: commonHandler,
		Foo:    fooHandler,
	}
}
