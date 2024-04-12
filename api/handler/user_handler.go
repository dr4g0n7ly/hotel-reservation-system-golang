package api

import (
	types "github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Foo",
	}
	return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Foo",
	}
	return c.JSON(u)
}
