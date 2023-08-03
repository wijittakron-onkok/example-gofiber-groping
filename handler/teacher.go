package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Teacher(c *fiber.Ctx) error {
	return c.SendString("Function teacher.")
}

func TeacherProfile(c *fiber.Ctx) error {
	return c.SendString("Function teacher profile.")
}

func TeacherSetting(c *fiber.Ctx) error {
	return c.SendString("Function teacher setting.")
}
