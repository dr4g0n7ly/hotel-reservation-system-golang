package api

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
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

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type GenericResponse struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func InvalidCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(GenericResponse{
		Type: "error",
		Msg:  "Invalid credentials",
	})
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return InvalidCredentials(c)
		}
		return err
	}

	if !types.ValidPassword(user.EncryptedPassword, params.Password) {
		return InvalidCredentials(c)
	}

	res := AuthResponse{
		User:  user,
		Token: CreateTokenFromUser(user),
	}

	return c.JSON(res)
}

func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expiresAt := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"expiresAt": expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("failed to sign token with secret: ", err)
	}
	return tokenStr
}
