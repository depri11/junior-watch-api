package test

import (
	"fmt"
	"log"

	"github.com/depri11/junior-watch-api/pkg/constants"
	"github.com/depri11/junior-watch-api/pkg/database"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type BaseTest struct {
	Log logger.Logger
	Cfg *config.Config
	Db  *sqlx.DB
}

func NewBaseTest() (*BaseTest, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	var db *sqlx.DB
	var source string
	switch cfg.Database.DBDriver {
	case constants.DBDriverPostgres:
		source = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.SSLMode)
		db, err = sqlx.Open(constants.DBDriverPostgres, source)
		if err != nil {
			appLogger.Fatal("cannot connect to Pgx", err)
		}
	case constants.DBDriverSqlite:
		source = fmt.Sprintf("%s.db", cfg.Database.Name)
		db, err = sqlx.Open(constants.DBDriverSqlite, source)
		if err != nil {
			appLogger.Fatal("cannot connect to SQLite3", err)
		}
	}

	if err := db.Ping(); err != nil {
		appLogger.Fatal(err)
		return nil, err
	}

	err = database.RunMigrate(&database.ConfigMigrate{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
		DBDriver: cfg.Database.DBDriver,
	}, db)
	if err != nil {
		return nil, err
	}

	return &BaseTest{appLogger, cfg, db}, nil
}
