package main

import (
	"webapp/go/controllers"
	"webapp/go/database"

	"github.com/gofiber/fiber/v2"
)


func setUpRoutes(app *fiber.App) {
	// app.Get("/hello", controllers.Hello)
	// app.Get("/hi", controllers.Hi)
	app.Post("/add", controllers.AddUser)
}
  


func main() {

	database.ConnectDB()

   app := fiber.New()

   setUpRoutes(app)

//    app.Get("/",func (c fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
//    })


   app.Listen(":3000")
}