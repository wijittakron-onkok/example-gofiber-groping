package main

import (
	"github.com/wijittakron-onkok/example-gofiber-groping/router"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	router.SetupRoutes(app)

	app.Listen(":8443")
}
