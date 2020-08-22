package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/uberballo/stockservice/cmd/db"
	"github.com/uberballo/stockservice/cmd/server"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	db.Ping()
	server.StartServer()
}
