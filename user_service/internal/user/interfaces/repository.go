package interfaces

import (
	"context"

	"github.com/depri11/junior-watch-api/user_service/internal/models"
)

type UserRepository interface {
	SaveUser(context.Context, *models.CreateUser) (string, error)
}
