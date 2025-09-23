package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/graphzc/go-clean-template/internal/domain/entities"
	userRepo "github.com/graphzc/go-clean-template/internal/repositories/user"
	"github.com/graphzc/go-clean-template/internal/utils/servererr"
	"github.com/graphzc/go-clean-template/internal/utils/timeutil"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, in *UserRegisterInput) error
}

type service struct {
	userRepo userRepo.Repository
}

// @WireSet("Service")
func NewService(userRepo userRepo.Repository) Service {
	return &service{}
}

// Register implements Service.
func (s *service) Register(ctx context.Context, in *UserRegisterInput) error {
	// Check user exitsting
	exitstingUser, err := s.userRepo.FindByEmail(ctx, in.Email)

	if err != nil {
		log.Ctx(ctx).
			Error().
			Err(err).
			Msg("Failed to check existion user")

		return servererr.NewError(
			servererr.ErrorCodeInternalServerError,
			"Failed to check existion user")
	}

	if exitstingUser != nil {
		log.Ctx(ctx).
			Warn().
			Str("email", in.Email).Msg("User with the same email already exists")

		return servererr.NewError(
			servererr.ErrorCodeConflict,
			"User with the same email already exists",
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(([]byte(in.Password)), bcrypt.DefaultCost)
	if err != nil {
		return servererr.NewError(
			servererr.ErrorCodeInternalServerError,
			"Failed to hash the password",
		)
	}

	newUser := &entities.User{
		ID:        uuid.NewString(),
		Email:     in.Email,
		Password:  string(hashedPassword),
		Name:      in.Name,
		CreatedAt: timeutil.BangkokNow(),
		UpdatedAt: timeutil.BangkokNow(),
	}

	err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		log.Ctx(ctx).
			Error().
			Err(err).
			Msg("Fail to create a user")

		return servererr.NewError(
			servererr.ErrorCodeInternalServerError,
			"Fail to create new user",
		)
	}

	return nil

}
