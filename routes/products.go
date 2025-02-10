package routes

import (
	"errors"

	"github.com/alexesp/Go_Fiber.git/database"
	"github.com/alexesp/Go_Fiber.git/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product)Product{
return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct (c *fiber.Ctx) error{
	var product models.Product

	err := c.BodyParser(&product)
		if err != nil{
			return c.Status(400).JSON(err.Error())
		
	}
	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error{
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseProducts := []Product{}

	for _, product := range products{
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error{
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0{
		return errors.New("Producto no existe")
	}
	return nil
}
func GetProduct (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}

	if err := findProduct(id, &product)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}
	if err := findProduct(id, &product)
	err != nil{
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct{
		Name string `json:"name"`
	    SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData)
	err != nil{
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}