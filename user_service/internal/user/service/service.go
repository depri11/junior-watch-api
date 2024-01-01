package service

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/google/uuid"
)

type UserService struct {
	log logger.Logger
}

func NewUserService(log logger.Logger) *UserService {
	return &UserService{log: log}
}

func (u *UserService) Register(ctx context.Context, user *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {
	return &userService.CreateUserResponse{UserID: uuid.NewString()}, nil
}
