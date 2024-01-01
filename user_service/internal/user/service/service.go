package service

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/internal/user/interfaces"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
)

type UserService struct {
	log  logger.Logger
	repo interfaces.UserRepository
}

func NewUserService(log logger.Logger, repo interfaces.UserRepository) *UserService {
	return &UserService{log, repo}
}

func (u *UserService) Register(ctx context.Context, user *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {
	id, err := u.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userService.CreateUserResponse{UserID: id}, nil
}
