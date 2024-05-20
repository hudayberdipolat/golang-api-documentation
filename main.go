package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

const serverPort = ":8000"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Golang and swagger API documentation",
		AppName:      "Golang and Swagger API documentation",
		BodyLimit:    10 * 1024 * 1024,
	})

	app.Get("/", indexHandler)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	err := app.Listen(serverPort)
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
	app.Listen(serverPort)
}

func indexHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "My name is Hudayberdi Polatov",
	})
}
