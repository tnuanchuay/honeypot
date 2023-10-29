package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App

func Init() {
	app = fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(logger.New())

	app.Use(loggerMiddleware)
}

func Get(path string, handler fiber.Handler) {
	app.Get(path, handler)
}

func Post(path string, handler fiber.Handler) {
	app.Post(path, handler)
}

func Run() {
	app.Listen(addr())
}
