package main

import (
	"api_fiber/database"
	"api_fiber/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

)

func main() {
	app := fiber.New()
	database.Connect()
	routes.Setup(app)

	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	var port string
	if port = os.Getenv("PORT"); port != "" {
		port = "8080"
	}

	app.Listen(":" + port)
}