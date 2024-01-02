package service

import (
	"context"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/models"
	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/google/uuid"
)

type UserService struct {
	log        logger.Logger
	cfg        *config.Config
	userClient userService.UserServiceClient
}

func NewUserService(log logger.Logger, cfg *config.Config, userClient userService.UserServiceClient) *UserService {
	return &UserService{log, cfg, userClient}
}

func (s *UserService) CreateUser(ctx context.Context, payload models.CreateUser) (*models.CreateUserResponse, error) {

	res, err := s.userClient.CreateUser(ctx, &userService.CreateUserRequest{
		Username: payload.Username,
		Email:    payload.Email,
		Phone:    payload.Phone,
		Name:     payload.Name,
		Address:  payload.Address,
	})
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(res.UserID)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{UserID: id}, nil
}
