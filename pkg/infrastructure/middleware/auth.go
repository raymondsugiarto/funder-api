package middleware

import (
	config "github.com/raymondsugiarto/funder-api/config"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/module/authentication"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
)

// Protected protect routes
func Protected() fiber.Handler {
	cfg := config.GetConfig()
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(cfg.Server.Rest.SecretKey)},
		ErrorHandler:   jwtError,
		SuccessHandler: SuccessHandler,
		Extractor:      extractors.FromAuthHeader("Bearer"),
	})
}

func jwtError(c fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func SuccessHandler(c fiber.Ctx) error {
	authenticationSvc := fiber.MustGetState[authentication.Service](c.App().State(), authentication.ServiceName)
	userSessionDto, err := authenticationSvc.GetSession(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": "error", "message": "Failed to get user session", "data": nil})
	}
	c.Locals(entity.UserSessionKey, userSessionDto)

	return c.Next()
}
