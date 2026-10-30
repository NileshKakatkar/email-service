package email

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp(cfg Config) *fiber.App {
	app := fiber.New(fiber.Config{AppName: "ShopCore Email Service", BodyLimit: 1 << 20})
	app.Use(recover.New())
	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok", "service": "email"}) })
	go RunConsumer(context.Background(), cfg)
	return app
}
