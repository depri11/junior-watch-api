package service

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/internal/models"
	"github.com/depri11/junior-watch-api/user_service/internal/user/interfaces"
)

type UserService struct {
	log  logger.Logger
	repo interfaces.UserRepository
}

func NewUserService(log logger.Logger, repo interfaces.UserRepository) *UserService {
	return &UserService{log, repo}
}

func (u *UserService) Register(ctx context.Context, user *models.CreateUser) (*models.CreateUserResponse, error) {
	id, err := u.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &models.CreateUserResponse{UserID: id}, nil
}
