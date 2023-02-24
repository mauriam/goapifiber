package server

import (
	"goapi/config"
	"goapi/pkg/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ServerRun() {
	app := fiber.New()

	handler.GetRoutes(app)

	configurations, err := config.GetConfig()
	if err != nil {
		log.Fatalln("error load settings ")
	}
	app.Listen(":" + configurations.Port)
}
