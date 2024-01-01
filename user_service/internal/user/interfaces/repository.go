package interfaces

import (
	"context"

	userService "github.com/depri11/junior-watch-api/user_service/proto"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *userService.CreateUserRequest) (string, error)
}
