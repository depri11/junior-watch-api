package commands

import (
	"context"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/pkg/logger"
)

type CreateUserCmdHandler interface {
	Handle(ctx context.Context, command *CreateUserCommand) error
}

type createProductHandler struct {
	log logger.Logger
	cfg *config.Config
}

func NewCreateUserHandler(log logger.Logger, cfg *config.Config) *createProductHandler {
	return &createProductHandler{log: log, cfg: cfg}
}

func (c *createProductHandler) Handle(ctx context.Context, command *CreateUserCommand) error {
	return nil
}
