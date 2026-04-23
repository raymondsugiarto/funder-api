package handler

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
)

func ChangePassword(service usercredential.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		input := new(entity.PasswordInputDto)
		if err := c.Bind().Body(input); err != nil {
			log.WithContext(c).Errorf("error body parser")
			return fiber.NewError(fiber.StatusBadRequest, "errorSignIn")
		}

		userSession := c.Locals(entity.UserSessionKey).(*entity.UserSessionDto)

		err := service.ChangePassword(c.Context(), userSession.ID, input.Password)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
