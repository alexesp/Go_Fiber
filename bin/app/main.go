package main

import (
	"log"

	"github.com/alexesp/Go_Fiber.git/database"
	"github.com/alexesp/Go_Fiber.git/routes"
	"github.com/gofiber/fiber/v2"
)


func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to may")
}

func setupRoutes(app *fiber.App){
	app.Get("/", welcome)
	app.Post("/users", routes.CreateUser)
}

func main(){
	database.ConnectDb()
	app := fiber.New()

	//app.Get("/", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}