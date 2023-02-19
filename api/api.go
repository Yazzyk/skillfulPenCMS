package api

import (
	"fmt"
	"github.com/Yazzyk/skillfulPenCMS/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	jsoniter "github.com/json-iterator/go"
)

var app *fiber.App

func InitAPI() {
	app = fiber.New(fiber.Config{
		AppName:     conf.Config.AppName,
		JSONDecoder: jsoniter.Unmarshal,
		JSONEncoder: jsoniter.Marshal,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{Format: "${method}${status}${url}\n"}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New())

	routers()

	app.Listen(fmt.Sprintf(":%d", conf.Config.Server.Port))
}
