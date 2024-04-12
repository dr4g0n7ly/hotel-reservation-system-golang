package main

import (
	api "github.com/dr4g0n7ly/hotel-management-system-golang/api/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(":5000")
}
