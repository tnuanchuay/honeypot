package pot

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tnuanchuay/honeypot/mysql"
	"net/http"
	"time"
)

var errDuplicate = errors.New("duplicate url")

func CreateGetHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		CreateHandler(ctx, http.MethodGet)
		return nil
	}
}

func CreatePostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		CreateHandler(ctx, http.MethodPost)
		return nil
	}
}

func CreateHandler(ctx *fiber.Ctx, method string) {
	var pot Pot
	var err error = nil

	switch method {
	case "GET":
		err = ctx.QueryParser(&pot)
	case "POST":
		err = ctx.BodyParser(&pot)
	}
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = pot.Validate()
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		log.Errorf("Some required fields in pot are missing")
		return
	}

	err = Create(pot)
	if err != nil {
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func Create(pot Pot) error {
	fmt.Println(pot)
	err := mysql.Execute(`INSERT INTO POT(path, redirect_to, user, create_date) VALUES(?, ?, ?, ?)`, pot.Path, pot.Redirect, pot.User, time.Now())
	if err != nil {
		return err
	}

	return nil
}
