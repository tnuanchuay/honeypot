package catch

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if errors.Is(err, fiber.ErrNotFound) {
			err = nil
			CreateCatchIfPotExists(ctx)
		}

		if err != nil {
			return err
		}

		return err
	}
}

func CreateCatchIfPotExists(ctx *fiber.Ctx) {
	path := ctx.Path()

}
