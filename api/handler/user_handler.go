package api

import "github.com/gofiber/fiber/v2"

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "user 1"})
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "user 1"})
}
