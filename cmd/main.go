package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found!")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world 👋")
	})

	address := ":" + portString
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(app.Listen(address))
}