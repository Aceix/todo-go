package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aceix/todo-api/views"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	if err := mgm.SetDefaultConfig(
		nil,
		"todos",
		options.Client().ApplyURI(connectionString),
	); err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := 3000
	app := fiber.New()

	views.RegisterViews(app)

	fmt.Printf("App listening on %v", port)
	app.Listen(port)
}
