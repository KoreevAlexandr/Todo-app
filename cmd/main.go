package main

import (
	"log"
	todo "main/Downloads/Using"
	"main/Downloads/Using/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http %s", err.Error())
	}
}
