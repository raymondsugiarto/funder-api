package middleware

import (
	"github.com/golang-jwt/jwt"
	config "github.com/raymondsugiarto/funder-api/config"
	shared "github.com/raymondsugiarto/funder-api/shared/context"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	cfg := config.GetConfig()
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(cfg.Server.Rest.SecretKey)},
		ErrorHandler:   jwtError,
		SuccessHandler: SuccessHandler,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func SuccessHandler(c *fiber.Ctx) error {
	token := c.Locals(shared.UserContextKey).(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userCredentialsData := new(shared.UserCredentialData)
	userCredentialsData.ID = claims["id"].(string)
	userCredentialsData.UserID = claims["uid"].(string)
	c.Locals(shared.UserCredentialDataKey, userCredentialsData)
	return c.Next()
}
