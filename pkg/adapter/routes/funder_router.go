package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/handler"
	"github.com/raymondsugiarto/funder-api/pkg/module/funder"
)

func FunderRouter(app *fiber.App, router fiber.Router) {
	svc := fiber.MustGetState[funder.Service](app.State(), funder.ServiceName)
	router.Post("/funders", handler.CreateFunder(svc))
	router.Get("/funders", handler.FindAllFunder(svc))
	router.Put("/funders/:id", handler.UpdateFunderByID(svc))
	router.Get("/funders/:id", handler.FindFunderByID(svc))
	router.Delete("/funders/:id", handler.DeleteFunderByID(svc))
}
