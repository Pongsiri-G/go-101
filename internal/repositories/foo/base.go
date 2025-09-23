package foo

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
}

type repository struct {
	db *sqlx.DB
}

// @WireSet("Repository")
func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
