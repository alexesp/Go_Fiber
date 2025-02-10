package routes

import (
	"errors"

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

func findUser(id int, user *models.User) error{
	database.Database.Db.Find(&user, "id = ?", id)

	if user.Id == 0{
		return errors.New("Usuario eno existe")
	}
	return nil
}

func GetUsers(c *fiber.Ctx) error{
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users{
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func GetUser (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}

	if err := findUser(id, &user)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}
	if err := findUser(id, &user)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct{
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData)
	err != nil{
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}

	if err := findUser(id, &user)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}

	err = database.Database.Db.Delete(&user).Error

	if err != nil{
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("El usuario borrado con exito.")
}