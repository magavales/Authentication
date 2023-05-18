package main

import (
	"JWT"
	"JWT/pkg/handler"
	"log"
)

func main() {
	router := new(handler.Handler).InitRouter()

	server := new(JWT.Server)
	err := server.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
