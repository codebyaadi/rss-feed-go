package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codebyaadi/rss-agg/internal/handlers"
	"github.com/codebyaadi/rss-agg/internal/helpers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found!")
	}

	http.HandleFunc("/", helpers.MethodMiddleware(http.MethodGet, handlers.ReadinessHandler))
	http.HandleFunc("/err", helpers.MethodMiddleware(http.MethodPost, handlers.ErrorHandler))

	address := ":" + portString
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}