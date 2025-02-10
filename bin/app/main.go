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
	//Punto final User
	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)
	//Punto final Producto
	app.Post("/products", routes.CreateProduct)
	app.Get("/products", routes.GetProducts)
	app.Get("/products/:id", routes.GetProduct)
	app.Put("/products/:id", routes.UpdateProduct)
}

func main(){
	database.ConnectDb()
	app := fiber.New()

	//app.Get("/", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}