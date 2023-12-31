package commands

import "github.com/depri11/junior-watch-api/api_gateway/internal/dto"

type UserCommands struct {
	CreateUser CreateUserCmdHandler
}

func NewUserCommands(createUser CreateUserCmdHandler) *UserCommands {
	return &UserCommands{CreateUser: createUser}
}

type CreateUserCommand struct {
	CreateDto *dto.CreateUserDto
}

func NewCreateUserCommand(createDto *dto.CreateUserDto) *CreateUserCommand {
	return &CreateUserCommand{CreateDto: createDto}
}
