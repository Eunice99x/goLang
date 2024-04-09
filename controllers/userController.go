package controllers

import (
	"webapp/go/database"
	"webapp/go/model"

	"github.com/gofiber/fiber/v2"
)

// func Hello( c fiber.Ctx) error {
// 	return c.SendString("Marhaba!")
// }

// func Hi(c fiber.Ctx) error {
// 	return c.SendString("Aslam Alikom!")
// }


func AddUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}