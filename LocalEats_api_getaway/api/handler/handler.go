package handler

import (
	"api_getaway/config"
	user "api_getaway/genproto/proto"
	"api_getaway/pkg"
	"api_getaway/pkg/logger"
	"log/slog"
)

type Handler struct {
	UserClient        user.AuthServiceClient
	Logger            *slog.Logger
}

func NewHandlerRepo() *Handler {
	cfg := config.Load()
	return &Handler{
		UserClient:        pkg.NewAuthServiceClient(cfg),
		Logger:            logger.NewLogger(),
	}
}
