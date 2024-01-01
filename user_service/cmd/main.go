package main

import (
	"flag"
	"log"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/pkg/postgres"
	"github.com/depri11/junior-watch-api/user_service/config"
	"github.com/depri11/junior-watch-api/user_service/internal/server"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("UserService")

	pgxConn, err := postgres.NewPgxConn(&postgres.PgConfig{Host: cfg.Postgres.Host, Port: cfg.Postgres.Port, User: cfg.Postgres.User, Password: cfg.Postgres.Password, Database: cfg.Postgres.Database, SSLMode: cfg.Postgres.SSLMode})
	if err != nil {
		appLogger.Fatal("cannot connect to postgres", err)
	}
	defer pgxConn.Close()
	appLogger.Infof("%-v", pgxConn.Stat())

	s := server.NewServer(appLogger, cfg, pgxConn)
	appLogger.Fatal(s.Run())
}
