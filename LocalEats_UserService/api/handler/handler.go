package handler

import (
	postgres "AuthService/Storage"
	"AuthService/pkg/logger"
	"database/sql"
	"log/slog"
)

type Handler struct {
	logger   *slog.Logger
	UserRepo *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		UserRepo: postgres.NewUserRepo(db),
		logger:   logger.NewLogger(),
	}
}
