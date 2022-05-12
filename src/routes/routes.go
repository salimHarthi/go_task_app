package routes

import (
	"go_task_app/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Get("/user/:id", controllers.GetUser)

	api.Get("/task/:id", controllers.GetallTasks)
	api.Put("/task/:id", controllers.UpdateTasks)
	api.Delete("/task/:id", controllers.DeleteTasks)
	api.Post("/task", controllers.CreateTasks)
}
