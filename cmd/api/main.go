package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/x87w/zenith/internal/api"
)

func main() {
	app := fiber.New()
	api.SetupRoutes(app)
	fmt.Println("API running on :8080")

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
