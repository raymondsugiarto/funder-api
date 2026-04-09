package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/middleware"
)

func InitRouter(app *fiber.App) {
	auth := app.Group("/api/auth")
	AuthRouter(app, auth)

	// TODO: middleware auth
	api := app.Group("/api", middleware.Protected())
	FunderRouter(app, api)
	ContractRouter(app, api)
	ContractPaymentRouter(app, api)
}
