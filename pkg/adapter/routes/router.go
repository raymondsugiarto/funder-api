package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/handler"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/middleware"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
)

func InitRouter(app *fiber.App) {
	auth := app.Group("/api/auth")
	AuthRouter(app, auth)

	// TODO: middleware auth
	api := app.Group("/api", middleware.Protected())

	userCredentialSvc := fiber.MustGetState[usercredential.Service](app.State(), usercredential.ServiceName)
	api.Put("user-credential/password", handler.ChangePassword(userCredentialSvc))

	FunderRouter(app, api)
	ContractRouter(app, api)
	ContractPaymentRouter(app, api)
}
