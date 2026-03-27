package server

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/raymondsugiarto/funder-api/config"
	"github.com/raymondsugiarto/funder-api/pkg/adapter/routes"
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

	app := fiber.New()

	app.Use(healthcheck.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	apiGroup := app.Group("/api")
	routes.InitRouter(apiGroup)

	err := app.Listen(":" + strconv.Itoa(cfg.Server.Rest.Port))
	if err != nil {
		log.Fatal(err)
	}
}
