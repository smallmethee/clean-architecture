package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(role string) func(*fiber.Ctx) error {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {

			user := ctx.Locals("user").(*jwt.Token)
			claims, ok := user.Claims.(jwt.MapClaims)
			if !ok {
				return ctx.
					Status(fiber.StatusUnauthorized).
					JSON(fiber.Map{
						"status":      fiber.ErrUnauthorized.Message,
						"status_code": fiber.StatusUnauthorized,
						"message":     "",
						"result":      nil,
					})
			}

			if claims["user_id"] != nil {
				return ctx.Next()
			}

			// roles := claims["roles"].([]interface{})

			// // handler role
			// for _, roleInterface := range roles {
			// 	roleMap := roleInterface.(map[string]interface{})
			// 	if roleMap["role"] == role {
			// 		return ctx.Next()
			// 	}
			// }

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"status":      fiber.ErrUnauthorized.Message,
					"status_code": fiber.StatusUnauthorized,
					"message":     "",
					"result":      nil,
				})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"status":      fiber.ErrBadRequest.Message,
						"status_code": fiber.StatusBadRequest,
						"message":     "",
						"result":      nil,
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(fiber.Map{
						"status":      fiber.ErrUnauthorized.Message,
						"status_code": fiber.StatusUnauthorized,
						"message":     "",
						"result":      nil,
					})
			}
		},
	})
}
