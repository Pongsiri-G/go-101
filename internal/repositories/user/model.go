package user

import (
	"time"
	"github.com/google/uuid"
)

type Model struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`

}
