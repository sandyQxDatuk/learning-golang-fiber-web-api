package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sandyQxDatuk/learning-golang-fiber-web-api/apps/common/databases"
)

func main() {
	databases.ConnectDatabase()
	app := fiber.New()
	app.Listen(":8000")
}
