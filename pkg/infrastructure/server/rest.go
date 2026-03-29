package server

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/raymondsugiarto/funder-api/config"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/routes"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/database"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/middleware"
	"github.com/raymondsugiarto/funder-api/pkg/infrastructure/middleware/organization"
	"github.com/raymondsugiarto/funder-api/shared/di"
)

type Rest struct {
}

func NewRest() *Rest {
	return &Rest{}
}

func (s *Rest) Initialize() {
	os.Setenv("TZ", "UTC")
	cfg := config.GetConfig()

	environment := cfg.Server.Rest.Env
	log.Infof("Environment: %s", environment)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.DefaultErrorHandler(),
	})
	initDatabase()

	app.Use(healthcheck.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	di.NewContainer(app).RegisterServices()

	app.Use(middleware.DefaultResponseHandler(), organization.New())

	routes.InitRouter(app)

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	err := app.Listen(":" + strconv.Itoa(cfg.Server.Rest.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase() {
	cfg := config.GetConfig()
	sqlConn, err := database.NewSQLConnection(cfg.Database.Main, cfg.Database.Main.Schema)
	if err != nil {
		log.Fatal(err)
	}
	database.DBConn = sqlConn.GetConn()
}
