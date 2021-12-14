package router

import (
	noteRouter "github.com/arunap2509/notes-api/internal/routes/note"
	userRouter "github.com/arunap2509/notes-api/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRouter(app *fiber.App) {

	api := app.Group("api", logger.New())

	noteRouter.SetupNotesRoute(api)

	userRouter.SetUpUserRoutes(api)
}
