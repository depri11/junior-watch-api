package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/grpc_client"
	v1 "github.com/depri11/junior-watch-api/api_gateway/internal/user/delivery/http/v1"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	"github.com/depri11/junior-watch-api/pkg/interceptors"
	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type server struct {
	gin *gin.Engine
	log logger.Logger
	cfg *config.Config
	v   *validator.Validate
	im  interceptors.InterceptorManager
	ps  *service.UserService
}

func NewServer(gin *gin.Engine, log logger.Logger, cfg *config.Config) *server {
	return &server{gin: gin, log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	userServiceConn, err := grpc_client.NewUserServiceConn(ctx, s.cfg, s.im, s.cfg.Grpc.UserServicePort)
	if err != nil {
		return err
	}

	defer userServiceConn.Close() // nolint: errcheck
	userClient := userService.NewUserServiceClient(userServiceConn)

	s.ps = service.NewUserService(s.log, s.cfg, userClient)

	usersHandlers := v1.NewUserHandlers(s.gin.Group(s.cfg.Http.UsersPath), s.log, s.cfg, s.v, s.ps)
	usersHandlers.Routes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	select {
	case v := <-quit:
		s.log.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.log.Errorf("ctx.Done: %v", done)
	}

	s.log.Info("Server Exited Properly")

	return nil
}
