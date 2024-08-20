package main

import (
	postgres "AuthService/Storage"
	"AuthService/api"
	"AuthService/config"
	pb "AuthService/genproto/users"
	"AuthService/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().USER_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	r := api.NewRouter(db)
	go r.Run(config.Load().USER_ROUTER)

	userservice := service.NewAuthServiceServer(db)
	service := grpc.NewServer()
	pb.RegisterAuthServiceServer(service, userservice)

	fmt.Printf("Server is listening on port %s\n", config.Load().USER_SERVICE)
	if err = service.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
