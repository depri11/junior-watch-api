package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/depri11/junior-watch-api/pkg/constants"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/config"
	"github.com/depri11/junior-watch-api/user_service/internal/server"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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

	var db *sqlx.DB

	switch cfg.Database.DBDriver {
	case constants.DBDriverPostgres:
		source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.SSLMode)
		db, err = sqlx.Open(constants.DBDriverPostgres, source)
		if err != nil {
			appLogger.Fatal("cannot connect to Pgx", err)
		}
		defer db.Close()
	case constants.DBDriverSqlite:
		source := fmt.Sprintf("%s.db", cfg.Database.Name)
		db, err = sqlx.Open(constants.DBDriverSqlite, source)
		if err != nil {
			appLogger.Fatal("cannot connect to SQLite3", err)
		}
		defer db.Close()
	}
	if err != nil {
		appLogger.Fatalf("cannot connect to %s", constants.DBDriverSqlite, err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	s := server.NewServer(appLogger, cfg, db)
	appLogger.Fatal(s.Run())
}
