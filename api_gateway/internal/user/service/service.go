package service

import (
	"context"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/models"
	"github.com/depri11/junior-watch-api/pkg/logger"
	go_proto "github.com/depri11/junior-watch-api/pkg/proto"
	"github.com/google/uuid"
)

type UserService struct {
	log        logger.Logger
	cfg        *config.Config
	userClient go_proto.UserServiceClient
}

func NewUserService(log logger.Logger, cfg *config.Config, userClient go_proto.UserServiceClient) *UserService {
	return &UserService{log, cfg, userClient}
}

func (s *UserService) CreateUser(ctx context.Context, payload models.CreateUser) (*models.CreateUserResponse, error) {

	res, err := s.userClient.Register(ctx, &go_proto.CreateUserRequest{
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
