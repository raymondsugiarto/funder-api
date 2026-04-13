package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/handler"
	"github.com/raymondsugiarto/funder-api/pkg/module/contract"
	contractpayment "github.com/raymondsugiarto/funder-api/pkg/module/contract/contract_payment"
)

func ContractRouter(app *fiber.App, router fiber.Router) {
	svc := fiber.MustGetState[contract.Service](app.State(), contract.ServiceName)
	router.Post("/contracts", handler.CreateContract(svc))
	router.Get("/contracts", handler.FindAllContract(svc))
	router.Get("/contracts/dashboard", handler.FindAllDashboard(svc))
	router.Get("/contracts/:id", handler.FindContractByID(svc))
	router.Put("/contracts/:id", handler.UpdateContractByID(svc))
	router.Delete("/contracts/:id", handler.DeleteContractByID(svc))
}

func ContractPaymentRouter(app *fiber.App, router fiber.Router) {
	svc := fiber.MustGetState[contractpayment.Service](app.State(), contractpayment.ServiceName)
	router.Post("/contracts/:id/payments", handler.CreateContractPayment(svc))

	router.Get("/contract-payments", handler.FindAllContractPayment(svc))
	router.Get("/contract-payments/:id", handler.FindContractPaymentByID(svc))
	router.Put("/contract-payments/:id", handler.UpdateContractPaymentByID(svc))
	router.Delete("/contract-payments/:id", handler.DeleteContractPaymentByID(svc))
}
