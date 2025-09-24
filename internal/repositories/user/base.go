package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/graphzc/go-clean-template/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
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

func (r *repository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `
		SELECT
			id, name, email, password, created_at, updated_at
		FROM
			users
		WHERE
			email = $1;`

	var userModel Model
	err := r.db.GetContext(ctx, &userModel, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return userModel.ToUserEntity(), nil
}

func (r *repository) Create(ctx context.Context, user *entities.User) error {
	model, err := FromUserEntity(user)
	if err != nil {
		return err
	}
	if model == nil {
		return ErrNullUser
	}
	query := `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES (:id, :name, :email, :password, :created_at, :updated_at);`

	result, err := r.db.NamedExecContext(ctx, query, model)
	if err != nil {
		return err
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return ErrNoRowsAffected
	}

	return nil
}
