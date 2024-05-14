package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/codebyaadi/rss-agg/internal/database"
	"github.com/codebyaadi/rss-agg/internal/handlers"
	"github.com/codebyaadi/rss-agg/internal/helpers"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found!")
	}

	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		log.Fatal("POSTGRES_URL not found!")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to Postgres database!", err)
	}

	queries := database.New(conn)

	apiCfg := handlers.ApiConfig{
		DB: queries,
	}

	http.HandleFunc("/", helpers.MethodMiddleware(http.MethodGet, handlers.ReadinessHandler))
	http.HandleFunc("/err", helpers.MethodMiddleware(http.MethodPost, handlers.ErrorHandler))
	http.HandleFunc("/user", helpers.MethodMiddleware(http.MethodPost, apiCfg.CreateUserHandler))

	address := ":" + portString
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}