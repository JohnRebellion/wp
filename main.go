package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Post("/api/webhook/staging", func(c *fiber.Ctx) error {
		requestBody := bytes.NewBuffer(c.Body())
		resp, err := http.Post(fmt.Sprintf("%s/api/webhook", os.Getenv("DOMAIN")), "application/json", requestBody)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		return c.Send(body)
	})
	app.Listen(":9999")
}
