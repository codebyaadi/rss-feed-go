package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/codebyaadi/rss-agg/internal/database"
	"github.com/codebyaadi/rss-agg/pkg/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on enviroment variables")
	}

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
		log.Fatalf("Can't connect to Postgres database! %v", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("Cannot reach postgres database %v", err)
	}

	queries := database.New(conn)

	apiCfg := handlers.ApiConfig{
		DB: queries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /root", handlers.ReadinessHandler)
	mux.HandleFunc("POST /err", handlers.ErrorHandler)
	mux.HandleFunc("POST /user", apiCfg.CreateUserHandler)

	address := ":" + portString
	server := &http.Server{
		Addr: address,
		Handler: mux,
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)

	go func() {
		log.Printf("Server is running on %s\n", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", address, err)
		}
	}()

	<-shutdownCh
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}