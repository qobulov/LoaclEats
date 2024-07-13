package main

import (
	postgres "AuthService/Storage"
	"AuthService/api"
	"AuthService/config"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	r := api.NewRouter(db)
	r.Run(config.Load().USER_ROUTER)
}
