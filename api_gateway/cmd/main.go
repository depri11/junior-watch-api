package main

import (
	"flag"
	"log"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/server"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("ApiGateway")

	gin := mux.NewRouter()

	s := server.NewServer(gin, appLogger, cfg)
	appLogger.Fatal(s.Run())
}
