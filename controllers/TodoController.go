package controllers

import (
	"github.com/Aceix/todo-api/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodosBy(ctx *fiber.Ctx) {
	coll := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	if err := coll.SimpleFind(&todos, bson.D{}); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}

func GetTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := models.Todo{}
	coll := mgm.Coll(&models.Todo{})

	if err := coll.FindByID(id, &todo); err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todo": todo,
	})
}

func CreateTodo(ctx *fiber.Ctx) {
	params := new(struct {
		Title string
		Desc  string
	})

	if err := ctx.BodyParser(&params); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid request body",
		})
		return
	}

	if len(params.Title) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "Title not specified",
		})
		return
	}

	todo := models.CreateTodo(
		params.Title,
		params.Desc,
	)
	if err := mgm.Coll(todo).Create(todo); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

func UpdateTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	params := new(struct {
		Title string
		Desc string
		Done bool
	})

	if err := ctx.BodyParser(&params); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	todo := &models.Todo{}
	coll := mgm.Coll(todo)

	// get todo
	if err := coll.FindByID(id, todo); err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found",
		})
		return
	}

	if len(params.Title) > 0 {
		todo.Title = params.Title
	}
	if len(params.Desc) > 0 {
		todo.Desc = params.Desc
	}
	todo.Done = params.Done

	if err := coll.Update(todo); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

func DeleteTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	coll := mgm.Coll(todo)

	err := coll.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	if err := coll.Delete(todo); err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
