package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/moodrets/go-tickets/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	server := http.NewServeMux()

	routes.TicketRoutes(server)
	routes.UserRoutes(server)

	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), server)
}
