package service

import (
	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/commands"
	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
)

type UserService struct {
	Commands *commands.UserCommands
}

func NewUserService(log logger.Logger, cfg *config.Config, userClient userService.UserServiceClient) *UserService {

	createUserHandler := commands.NewCreateUserHandler(log, cfg)

	userCommands := commands.NewUserCommands(createUserHandler)

	return &UserService{Commands: userCommands}
}
