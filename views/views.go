package views

import (
	"github.com/Aceix/todo-api/controllers"
	"github.com/gofiber/fiber"
)

func RegisterViews(app *fiber.App) {
	// create
	app.Post("/api/v1/todos", controllers.CreateTodo)
	
	// bulk read
	app.Get("/api/v1/todos", controllers.GetTodosBy)

	// detail read
	app.Get("/api/v1/todos/:id", controllers.GetTodo)
	
	// update
	app.Patch("/api/v1/todos/:id", controllers.UpdateTodo)
	
	// delete
	app.Delete("/api/v1/todos/:id", controllers.DeleteTodo)
}
