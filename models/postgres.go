package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/zero-one-group/fullstack-go/config"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     config.Env.Database.Host,
		Port:     config.Env.Database.Port,
		User:     config.Env.Database.Username,
		Password: config.Env.Database.Password,
		Database: config.Env.Database.Name,
		SSLMode:  config.Env.Database.SSLMode,
	}
}

// Open will open a SQL connection with the provided
// Postgres database. Callers of Open need to ensure
// the connection is eventually closed via the
// db.Close() method.
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
