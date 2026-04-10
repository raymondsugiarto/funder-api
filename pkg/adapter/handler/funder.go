package handler

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	"github.com/raymondsugiarto/funder-api/pkg/model"
	"github.com/raymondsugiarto/funder-api/pkg/module/funder"
)

func CreateFunder(service funder.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(entity.FunderRequest)

		if err := c.Bind().Body(request); err != nil {
			log.WithContext(c).Errorf("error body parser")
			return fiber.NewError(fiber.StatusBadRequest, "errorSignIn")
		}

		response, err := service.Create(c, request.ToDto())
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func FindFunderByID(service funder.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")

		response, err := service.FindByID(c, id)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func FindAllFunder(service funder.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		query := new(entity.FunderFilterDto)
		if err := c.Bind().Query(query); err != nil {
			log.WithContext(c).Errorf("error query parser", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		userSession := c.Locals(entity.UserSessionKey).(*entity.UserSessionDto)
		user := userSession.UserCredential.User
		if user.UserType == model.FUNDER {
			funder, err := service.FindByUserID(c, user.ID)
			if err != nil {
				log.WithContext(c).Errorf("error find funder by user id", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			query.FunderIDParent = funder.ID
		}

		response, err := service.FindAll(c, query)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}
