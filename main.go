package main

import (
	"webapp/go/database"

	"github.com/gofiber/fiber/v3"
)

func main() {

	database.ConnectDB()
   app := fiber.New()

   app.Get("/",func (c fiber.Ctx) error {
		return c.SendString("Hello, World!")
   })

   app.Listen(":3000")
}