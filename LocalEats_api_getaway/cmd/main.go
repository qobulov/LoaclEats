package main

import (
	"api_getaway/api"
	"api_getaway/config"
)

func main() {

	r := api.NewRouter()
	r.Run(config.Load().API_GATEWAY)
}
