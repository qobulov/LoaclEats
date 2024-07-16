package postgres

import (
	"database/sql"
	"fmt"
	"Orders/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
