package rest

import (
	"main/internal/services"

	"github.com/gofiber/fiber/v2"
)

func Handler() {
	app := fiber.New()

	app.Post("/sign-up", services.SignUp)

	app.Listen(":8080")
}
