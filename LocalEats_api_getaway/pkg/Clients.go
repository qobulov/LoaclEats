package pkg

import (
	"api_getaway/config"
	"api_getaway/genproto/orders"
	user "api_getaway/genproto/users"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthServiceClient(cfg *config.Config) user.AuthServiceClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return user.NewAuthServiceClient(conn)
}

func NewOrderServiceClient(cfg *config.Config) orders.OrderServiceClient {
	conn, err := grpc.NewClient(cfg.ORDER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return orders.NewOrderServiceClient(conn)
}
