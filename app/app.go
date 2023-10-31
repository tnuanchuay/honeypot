package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tnuanchuay/honeypot/honeypot/catch"
	"github.com/tnuanchuay/honeypot/honeypot/pot"
)

var app *fiber.App

func Init() {
	app = fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(logger.New())
	app.Use(catch.Handler())

	app.Get("/honey/new", pot.CreateGetHandler())
	app.Post("/honey/new", pot.CreatePostHandler())
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
