package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/depri11/junior-watch-api/pkg/logger"
	go_proto "github.com/depri11/junior-watch-api/pkg/proto"
	"github.com/depri11/junior-watch-api/user_service/config"
	"github.com/depri11/junior-watch-api/user_service/internal/user/delivery"
	"github.com/depri11/junior-watch-api/user_service/internal/user/repository"
	"github.com/depri11/junior-watch-api/user_service/internal/user/service"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	log logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewServer(logger logger.Logger, cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{log: logger, cfg: cfg, db: db}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l, err := net.Listen("tcp", s.cfg.GRPCServer.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.GRPCServer.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.GRPCServer.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.GRPCServer.MaxConnectionAge * time.Minute,
		Time:              s.cfg.GRPCServer.Timeout * time.Minute,
	}))

	repoUser := repository.NewUserRepository(&s.log, s.db)
	serviceUser := service.NewUserService(&s.log, repoUser)

	deliveryUser := delivery.NewUserDelivery(serviceUser, s.log)
	go_proto.RegisterUserServiceServer(server, deliveryUser)

	go func() {
		s.log.Infof("GRPC Server is listening on port: %v", s.cfg.GRPCServer.Port)
		s.log.Fatal(server.Serve(l))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.log.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.log.Errorf("ctx.Done: %v", done)
	}

	s.log.Info("Server Exited Properly")

	return nil
}
