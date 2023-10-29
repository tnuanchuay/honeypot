package honeypot

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/tnuanchuay/honeypot/mysql"
	"net/http"
)

var errDuplicate = errors.New("duplicate url")

func CreateGetHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var pot Pot
		err := ctx.QueryParser(&pot)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			goto done
		}

		err = Create(pot)

		if errors.Is(err, errDuplicate) {
			ctx.Status(http.StatusConflict)
			goto done
		}

		ctx.Status(http.StatusCreated)

	done:
		return nil
	}
}

func CreatePostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var pot Pot
		err := ctx.BodyParser(&pot)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			goto done
		}

		err = Create(pot)

		if errors.Is(err, errDuplicate) {
			ctx.Status(http.StatusConflict)
			goto done
		}

		ctx.Status(http.StatusCreated)

	done:
		return nil
	}
}

func Create(pot Pot) error {
	err := mysql.Execute(`INSERT INTO POT(PATH, REDIRECT, USER) VALUES(?, ?, ?)`, pot.Path, pot.Redirect, pot.User)
	if err != nil {
		return err
	}

	return nil
}

func createTablePot() {
	err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS honey (
		id 					INT 			PRIMARY KEY NOT NULL 	AUTO_INCREMENT,
		redirect_to 		VARCHAR(1000) 				NOT NULL,
		user_id 			VARCHAR(100) 				NOT NULL,
		create_date 		TIMESTAMP 					NOT NULL,
		url 				VARCHAR(200)
	);
`)

	if err != nil {
		panic(err)
	}
}
