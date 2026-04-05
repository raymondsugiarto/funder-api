package routes

import "github.com/gofiber/fiber/v3"

func InitRouter(app *fiber.App) {
	auth := app.Group("/api/auth")
	AuthRouter(app, auth)

	// TODO: middleware auth
	api := app.Group("/api")
	FunderRouter(app, api)
	ContractRouter(app, api)
	ContractPaymentRouter(app, api)
}
