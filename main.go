package main

import (
	"example/testapp/controllers"
	"example/testapp/initializers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	fmt.Println("init main")
	initializers.ConnectToDB()
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("hello")
		return c.SendString("Hello, World!")
	})

	app.Post("/create", controllers.UserCreate)
	app.Post("/user", controllers.GetUserById)
	app.Get("/allusers", controllers.GetAllUsers)
	app.Get("/search", controllers.GetSynonyms)
	app.Put("/update", controllers.UpDateUser)
	app.Delete("/delete", controllers.Delete)
}

func main() {
	app := fiber.New()

	setUpRoutes(app)
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
