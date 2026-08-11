package api

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":   "Zenith API",
			"status": "online",
		})
	})

	app.Get("/api/vms", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"vms": []string{},
		})
	})

	app.Post("/api/register", Register)

	app.Post("/api/vms/create", CreateVM)
}
