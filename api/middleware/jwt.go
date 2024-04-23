package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT Auth")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		fmt.Println("did not get token header")
		return fmt.Errorf("unauthorized")
	}

	claims, err := ValidateToken(token[0])
	if err != nil {
		return err
	}

	fmt.Println(claims)

	// Check token expiration
	expFloat := claims["expiresAt"].(float64)
	expiration := int64(expFloat)
	if time.Now().Unix() > expiration {
		fmt.Println("token has expired")
		return fmt.Errorf("unauthorized")
	}

	return c.Next()
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("parsing error: ", err)
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("error type claim")
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
