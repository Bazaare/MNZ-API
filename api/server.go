package api

import (
	"MNZ/api/controllers"
	"github.com/joho/godotenv"
	"log"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	server.Initialise()
	server.Run(":8080")
}
