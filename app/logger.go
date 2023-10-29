package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnuanchuay/honeypot/log"
)

var loggerMiddleware = func(ctx *fiber.Ctx) error {
	log.Info(ctx.BaseURL())

	return ctx.Next()
}
