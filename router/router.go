package router

import (
	"github.com/wijittakron-onkok/example-gofiber-groping/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	student := v1.Group("/student")
	student.Get("/", handler.Student)
	student.Get("/profile", handler.StudentProfile)
	student.Get("/setting", handler.StudentSetting)

	teacher := v1.Group("/teacher")
	teacher.Get("/", handler.Teacher)
	teacher.Get("/profile", handler.TeacherProfile)
	teacher.Get("/setting", handler.TeacherSetting)
}
