package middleware

import (
	"fmt"

	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		fmt.Println("not ok")
		return fmt.Errorf("not authorized")
	}
	if !user.IsAdmin {
		fmt.Println("user not admin")
		fmt.Printf("%+v", user)
		return fmt.Errorf("not authorized")
	}
	return c.Next()
}
