package main

import (
	"golang-test/database"
	"golang-test/todo"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":9000"))
}
