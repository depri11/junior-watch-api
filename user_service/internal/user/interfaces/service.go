package interfaces

import (
	"context"

	userService "github.com/depri11/junior-watch-api/user_service/proto"
)

type UserService interface {
	Register(ctx context.Context, user *userService.CreateUserRequest) (*userService.CreateUserResponse, error)
}
