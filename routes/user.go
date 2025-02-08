package routes

import (
	"github.com/alexesp/Go_Fiber.git/database"
	"github.com/alexesp/Go_Fiber.git/models"
	"github.com/gofiber/fiber/v2"
)

type User struct{
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User{
	return User{ID: userModel.Id, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error{
	var user models.User

	if err :=  c.BodyParser(&user)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}