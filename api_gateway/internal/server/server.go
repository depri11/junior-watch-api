package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	grpcClient "github.com/depri11/junior-watch-api/api_gateway/internal/grpc_client"
	v1 "github.com/depri11/junior-watch-api/api_gateway/internal/user/delivery/http/v1"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	"github.com/depri11/junior-watch-api/pkg/interceptors"
	"github.com/depri11/junior-watch-api/pkg/logger"
	go_proto "github.com/depri11/junior-watch-api/pkg/proto"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type server struct {
	mux *mux.Router
	log logger.Logger
	cfg *config.Config
	v   *validator.Validate
	im  interceptors.InterceptorManager
	ps  *service.UserService
}

func NewServer(mux *mux.Router, log logger.Logger, cfg *config.Config) *server {
	return &server{mux: mux, log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	userServiceConn, err := grpcClient.NewUserServiceConn(ctx, s.cfg, s.im, s.cfg.Grpc.UserServicePort)
	if err != nil {
		return err
	}

	defer userServiceConn.Close() // nolint: errcheck
	userClient := go_proto.NewUserServiceClient(userServiceConn)

	s.ps = service.NewUserService(s.log, s.cfg, userClient)

	usersHandlers := v1.NewUserHandlers(s.mux.PathPrefix(s.cfg.Http.UsersPath).Subrouter(), s.log, s.cfg, s.v, s.ps)
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
