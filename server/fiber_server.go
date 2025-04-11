package server

import (
	"fmt"

	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type (
	Server interface {
		Start()
	}

	fiberServer struct {
		app *fiber.App
		conf *conf.Config
	}
)

func NewFiberServer(conf *conf.Config) Server {
	fiberApp := fiber.New(fiber.Config{
		ReadBufferSize: 60 * 1024,
		DisableStartupMessage: false,
		AppName: conf.App.NAME,
	});

		return &fiberServer{
			app: fiberApp,
			conf: conf,
		}
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New());

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type",
	}));

	s.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "api is available",
		})
	});

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port);

	s.app.Listen(serverUrl);
}