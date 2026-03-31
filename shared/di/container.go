package di

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/database"
	"github.com/raymondsugiarto/funder-api/pkg/module/authentication"
	"github.com/raymondsugiarto/funder-api/pkg/module/funder"
	"github.com/raymondsugiarto/funder-api/pkg/module/user"
	usercredential "github.com/raymondsugiarto/funder-api/pkg/module/user-credential"
	"github.com/raymondsugiarto/funder-api/shared/database/transaction"
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
	c.gormManager()

	c.userCredentialService()
	c.userService()
	c.funderService()

	c.authenticationService()
}

func (c *container) gormManager() {
	gormTxManager := transaction.NewGormManager(database.DBConn)
	c.add(transaction.GormServiceName, gormTxManager)
}

func (c *container) gormManagerWithTx() transaction.Manager {
	return fiber.MustGetState[transaction.Manager](c.app.State(), transaction.GormServiceName)
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

func (c *container) userService() {
	userCredentialService := fiber.MustGetState[usercredential.Service](c.app.State(), usercredential.ServiceName)
	userRepository := user.NewRepository(database.DBConn)
	userService := user.NewService(userRepository, userCredentialService)
	c.add(user.ServiceName, userService)
}

func (c *container) funderService() {
	userService := fiber.MustGetState[user.Service](c.app.State(), user.ServiceName)
	funderRepository := funder.NewRepository(database.DBConn)
	funderService := funder.NewService(c.gormManagerWithTx(), funderRepository, userService)
	c.add(funder.ServiceName, funderService)
}

func (c *container) add(name string, service any) {
	if c.app.State().Has(name) {
		panic("Service already registered: " + name)
	}
	c.app.State().Set(name, service)
}
