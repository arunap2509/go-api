package main

import (
	"encoding/json"

	"github.com/arunap2509/notes-api/database"
	"github.com/arunap2509/notes-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	database.ConnectDB()

	router.SetUpRouter(app)

	app.Listen(":3000")
}
