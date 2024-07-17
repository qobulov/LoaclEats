package main

import (
	postgres "Orders/Storage"
	"Orders/config"
	pb "Orders/genproto/orders"
	"Orders/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().ORDER_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	userservice := service.NewOrderService(db)
	servic := grpc.NewServer()
	pb.RegisterOrderServiceServer(servic, userservice)

	fmt.Printf("Server is listening on port %s\n", config.Load().ORDER_SERVICE)
	if err = servic.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
