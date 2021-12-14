package userRoutes

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(router fiber.Router) {

	user := router.Group("users")

	user.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "response from the user router"})
	})
}
