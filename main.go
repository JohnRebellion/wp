package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Post("payment_requests", func(c *fiber.Ctx) error {
		log.Printf("Caller IP: %s\n", c.IPs())
		requestBody := bytes.NewBuffer(c.Body())
		client := &http.Client{}
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment_requests", os.Getenv("DOMAIN")), requestBody)
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(os.Getenv("XENDIT_SECRET_API_KEY"), "")
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		log.Println(string(body))

		return c.Send(body)
	})

	app.Post("/payment_requests", func(c *fiber.Ctx) error {
		log.Printf("Caller IP: %s\n", c.IPs())
		requestBody := bytes.NewBuffer(c.Body())
		resp, err := http.Post(fmt.Sprintf("%s/payment_requests", os.Getenv("DOMAIN")), "application/json", requestBody)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		return c.Send(body)
	})
	app.Listen(":9999")
}
