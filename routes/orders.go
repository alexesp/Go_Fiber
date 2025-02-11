package routes

import (
	"errors"
	"time"

	"github.com/alexesp/Go_Fiber.git/database"
	"github.com/alexesp/Go_Fiber.git/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct{
	ID uint `json:"id"`
	User User `json:"user"`
	Product Product `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user User, product Product)Order{
return Order{ID: orderModel.ID, User: user, Product: product, CreatedAt: orderModel.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error{
	var order models.Order

	err := c.BodyParser(&order)
	if err != nil{
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	err = findUser(order.UserRefer, &user)

	if err != nil{
            return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	err = findProduct(order.ProductRefer, &product)

	if err != nil{
            return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error{
	orders := []models.Order{}

	database.Database.Db.Find(&orders)

	responseOrders := []Order{}

	for _, order := range orders{
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	

	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order)error{
	database.Database.Db.Find(&order, "id = ?", id) 
	if order.ID == 0{
		return errors.New("Pedido no encontrado")
	}

return nil

}
func GetOrder(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil{
		return c.Status(400).JSON("Asegurase que id es de tipo entero")
	}

	err = findOrder(id, &order)
	if err != nil{
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}