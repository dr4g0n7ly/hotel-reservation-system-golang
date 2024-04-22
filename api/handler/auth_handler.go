package api

import (
	"fmt"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(store db.Store) *AuthHandler {
	return &AuthHandler{
		userStore: store.User,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var AuthParams AuthParams
	if err := c.BodyParser(&AuthParams); err != nil {
		return err
	}

	fmt.Println(AuthParams)

	return nil
}
