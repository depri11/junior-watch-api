package delivery

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/internal/user/interfaces"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
)

type UserDelivery struct {
	userService interfaces.UserService
	logger      logger.Logger
}

func NewUserService(userService interfaces.UserService, logger logger.Logger) *UserDelivery {
	return &UserDelivery{userService: userService, logger: logger}
}

func (u *UserDelivery) CreateUser(ctx context.Context, r *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {
	res, err := u.userService.Register(ctx, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}
