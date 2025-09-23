package foo

import (
	"github.com/graphzc/go-clean-template/internal/config"
	"github.com/graphzc/go-clean-template/internal/repositories/foo"
)

type Service interface {
}

type service struct {
	config  *config.Config
	fooRepo foo.Repository
}

// @WireSet("Service")
func NewService(
	config *config.Config,
	fooRepo foo.Repository,
) Service {
	return &service{
		config:  config,
		fooRepo: fooRepo,
	}
}
