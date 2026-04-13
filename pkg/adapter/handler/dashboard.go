package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/module/contract"
)

func FindAllDashboard(service contract.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		funderID := ""
		funderSession := c.Locals(entity.FunderSessionKey)
		if funderSession != nil {
			funderID = funderSession.(*entity.FunderDto).ID
		}

		response, err := service.ViewDashboard(c, funderID)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}
