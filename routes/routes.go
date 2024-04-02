package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Henrycall/learn-fiber/handlers"
)

func Routes (app *fiber.App) {
	app.Get("/" , handlers.ReadData)
	app.Get("/:id", handlers.ReadDataById)
	app.Post("/", handlers.InsertData)
	app.Delete("/:id", handlers.DeleteData)
	app.Patch("/:id", handlers.PatchData)
}