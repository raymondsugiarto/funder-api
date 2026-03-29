package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/handler"
	"github.com/raymondsugiarto/funder-api/pkg/module/authentication"
)

func AuthRouter(app *fiber.App, router fiber.Router) {
	authSvc := fiber.MustGetState[authentication.Service](app.State(), authentication.ServiceName)
	router.Post("/sign-in", handler.SignIn(authSvc))
}
