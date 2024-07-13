package pkg

import (
	"api_getaway/config"
	user "api_getaway/genproto/proto"
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
