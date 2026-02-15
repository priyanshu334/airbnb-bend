package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/airbnbbend/internal/config"
)

func main() {
	cfg := config.Load()
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Listen(":" + cfg.AppPort)

}
