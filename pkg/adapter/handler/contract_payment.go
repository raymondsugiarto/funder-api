package handler

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/entity"
	contractpayment "github.com/raymondsugiarto/funder-api/pkg/module/contract/contract_payment"
)

func CreateContractPayment(service contractpayment.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(entity.ContractPaymentRequest)

		if err := c.Bind().Body(request); err != nil {
			log.WithContext(c).Errorf("error body parser")
			return fiber.NewError(fiber.StatusBadRequest, "errorParseRequest")
		}

		attachmentUrl := `./storage/` + request.AttachmentFile.Filename
		err := c.SaveFile(request.AttachmentFile, attachmentUrl)
		if err != nil {
			log.WithContext(c).Errorf("error save file", err)
			return fiber.NewError(fiber.StatusInternalServerError, "errorSaveFile")
		}

		response, err := service.Create(c, request.ToDto(attachmentUrl))
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func FindContractPaymentByID(service contractpayment.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")

		response, err := service.FindByID(c, id)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func FindAllContractPayment(service contractpayment.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		query := new(entity.ContractPaymentFilterDto)
		if err := c.Bind().Query(query); err != nil {
			log.WithContext(c).Errorf("error query parser", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		funderSession := c.Locals(entity.FunderSessionKey)
		if funderSession != nil {
			query.FunderID = funderSession.(*entity.FunderDto).ID
		}

		response, err := service.FindAll(c, query)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func UpdateContractPaymentByID(service contractpayment.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(entity.ContractPaymentRequest)
		id := c.Params("id")

		if err := c.Bind().Body(request); err != nil {
			log.WithContext(c).Errorf("error body parser")
			return fiber.NewError(fiber.StatusBadRequest, "errorSignIn")
		}

		attachmentUrl := ""
		if request.AttachmentFile != nil {
			attachmentUrl := `./storage/` + request.AttachmentFile.Filename
			err := c.SaveFile(request.AttachmentFile, attachmentUrl)
			if err != nil {
				log.WithContext(c).Errorf("error save file", err)
				return fiber.NewError(fiber.StatusInternalServerError, "errorSaveFile")
			}
		}

		dto := request.ToDto(attachmentUrl)
		dto.ID = id

		response, err := service.Update(c, dto)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func DeleteContractPaymentByID(service contractpayment.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")

		err := service.Delete(c, id)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
