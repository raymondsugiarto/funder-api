package di

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/database"
	"github.com/raymondsugiarto/funder-api/pkg/module/authentication"
	"github.com/raymondsugiarto/funder-api/pkg/module/funder"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
)

type Container interface {
	RegisterServices()
}

type container struct {
	app *fiber.App
}

func NewContainer(app *fiber.App) Container {
	return &container{
		app: app,
	}
}

func (c *container) RegisterServices() {
	c.userCredentialService()
	c.funderService()

	c.authenticationService()
}

func (c *container) authenticationService() {
	userCredentialService := fiber.MustGetState[usercredential.Service](c.app.State(), usercredential.ServiceName)
	authenticationService := authentication.NewService(userCredentialService)
	c.add(authentication.ServiceName, authenticationService)
}

func (c *container) userCredentialService() {
	userCredentialRepository := usercredential.NewRepository(database.DBConn)
	userCredentialService := usercredential.NewService(userCredentialRepository)
	c.add(usercredential.ServiceName, userCredentialService)
}

func (c *container) funderService() {
	funderRepository := funder.NewRepository(database.DBConn)
	funderService := funder.NewService(funderRepository)
	c.add(funder.ServiceName, funderService)
}

func (c *container) add(name string, service any) {
	if c.app.State().Has(name) {
		panic("Service already registered: " + name)
	}
	c.app.State().Set(name, service)
}
