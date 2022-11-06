package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"main/Database"
	"main/Routes"
)

func main() {
	Database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	Routes.SetUp(app)
	app.Listen(":8000")
	//fmt.Println("Successfully Connected")

}
