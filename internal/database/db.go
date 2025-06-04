package database

import (
	"database/sql"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка ping: %w", err)
	}
	return db, nil
}
