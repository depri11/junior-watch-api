package database

import (
	"os"
	"path/filepath"

	"github.com/depri11/junior-watch-api/pkg/constants"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
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

func RunMigrate(cfg *ConfigMigrate, db *sqlx.DB) error {

	getwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "os.Getwd")
	}

	var driver database.Driver

	switch cfg.DBDriver {
	case constants.DBDriverPostgres:
		driver, _ = postgres.WithInstance(db.DB, &postgres.Config{})
	case constants.DBDriverSqlite:
		driver, _ = sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	default:
		return errors.New("unsupported database driver")
	}

	configPath := "file://" + filepath.Join(getwd, "..", "migrations")
	m, err := migrate.NewWithDatabaseInstance(
		configPath,
		db.DriverName(),
		driver,
	)
	if err != nil {
		return errors.Wrap(err, "migrations")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
