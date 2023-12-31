package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/client"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	"github.com/depri11/junior-watch-api/pkg/interceptors"
	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/go-playground/validator"
)

type server struct {
	log logger.Logger
	cfg *config.Config
	v   *validator.Validate
	im  interceptors.InterceptorManager
	ps  *service.UserService
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	userServiceConn, err := client.NewReaderServiceConn(ctx, s.cfg, s.im)
	if err != nil {
		return err
	}

	defer userServiceConn.Close() // nolint: errcheck
	userClient := userService.NewUserServiceClient(userServiceConn)

	s.ps = service.NewUserService(s.log, s.cfg, userClient)

	return nil
}
