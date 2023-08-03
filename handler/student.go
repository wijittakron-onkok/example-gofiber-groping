package handler

import (
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Username string `json:"username"`
	ID       int    `json:"student_id"`
}

func Student(c *fiber.Ctx) error {
	u := user{
		Username: "somchai",
		ID:       1234,
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "First data students.",
		"data":    u,
	})
}

func StudentProfile(c *fiber.Ctx) error {
	return c.SendString("Function student profile.")
}

func StudentSetting(c *fiber.Ctx) error {
	return c.SendString("Function student setting.")
}
