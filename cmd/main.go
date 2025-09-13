package main

import (
	"go-web-crud/internal/database"
	"go-web-crud/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET_KEY"))

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	database.InitDatabase()

	routes.InitRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
