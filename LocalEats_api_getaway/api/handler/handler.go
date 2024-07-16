package handler

import (
	"api_getaway/config"
	"api_getaway/genproto/orders"
	user "api_getaway/genproto/users"
	"api_getaway/pkg"
	"api_getaway/pkg/logger"
	"log/slog"
)

type Handler struct {
	UserClient  user.AuthServiceClient
	OrderClient orders.OrderServiceClient
	Logger      *slog.Logger
}

func NewHandlerRepo() *Handler {
	cfg := config.Load()
	return &Handler{
		UserClient:  pkg.NewAuthServiceClient(cfg),
		Logger:      logger.NewLogger(),
		OrderClient: pkg.NewOrderServiceClient(cfg),
	}
}
