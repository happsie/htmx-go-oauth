package repository

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/happsie/gohtmx/config"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

func NewDatabase(log *slog.Logger, config config.Config) (*sqlx.DB, error) {
	log.Info("connecting to database")
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@/%s", config.Database.User, config.Database.Password, config.Database.Database))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	log.Info("database connection established")
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}
	return db, nil
}
