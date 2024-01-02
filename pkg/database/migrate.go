package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/depri11/junior-watch-api/pkg/constants"
	"github.com/depri11/junior-watch-api/pkg/logger"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattes/migrate/source/file"
	"github.com/pkg/errors"
)

type ConfigMigrate struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	DBDriver string
}

func RunMigrate(cfg *ConfigMigrate, log logger.Logger) error {

	getwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "os.Getwd")
	}

	var dbUrl string

	switch cfg.DBDriver {
	case constants.DBDriverPostgres:
		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
	case constants.DBDriverSqlite:
		dbUrl = fmt.Sprintf("sqlite3://%s/%s.db", getwd, cfg.Name)
	default:
		return errors.New("unsupported database driver")
	}

	configPath := "file://" + filepath.Join(getwd, "..", "migrations")
	m, err := migrate.New(
		configPath,
		dbUrl,
	)
	if err != nil {
		return errors.Wrap(err, "migrations")
	}

	// remove all evrything
	err = m.Drop()
	if err != nil {
		return err
	}
	m.Close()

	// do new migration
	m, err = migrate.New(
		configPath,
		dbUrl,
	)
	if err != nil {
		return errors.Wrap(err, "migrations")
	}

	log.Infof("created migration path: %s", configPath)
	err = m.Up()
	if err != nil {
		return err
	}
	m.Close()

	return nil
}
