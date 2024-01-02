package interfaces

import (
	"context"

	"github.com/depri11/junior-watch-api/user_service/internal/models"
)

type UserService interface {
	Register(ctx context.Context, user *models.CreateUser) (*models.CreateUserResponse, error)
}
