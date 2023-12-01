package catch

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tnuanchuay/honeypot/honeypot/ipgeo"
	"github.com/tnuanchuay/honeypot/honeypot/pot"
	"github.com/tnuanchuay/honeypot/mysql"
	"net/http"
	"strings"
	"time"
)

func Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil && strings.HasPrefix(err.Error(), "Cannot GET") {
			err = nil
			CreateCatchIfPotExists(ctx)
		}

		if err != nil {
			log.Error(err)
		}

		return err
	}
}

func CreateCatchIfPotExists(ctx *fiber.Ctx) {
	path := ctx.Path()
	pot, err := pot.FindPotByPath(path)
	if err != nil {
		log.Error(err)
		return
	}

	if pot == nil {
		return
	}

	ctx.Redirect(fmt.Sprintf("https://%s", pot.Redirect), http.StatusFound)
	referer := ctx.Get("Referer")
	c := Catch{
		PotId:      pot.Id,
		RemoteIp:   ctx.IP(),
		Referer:    referer,
		CreateDate: time.Now(),
	}

	go Create(&c)
}

func Create(catch *Catch) {
	result, _ := mysql.Execute(`INSERT INTO CATCH(pot_id, remote_ip, referer, create_at) VALUES (?, ?, ?, ?)`, catch.PotId, catch.RemoteIp, catch.Referer, catch.CreateDate)
	lastId, _ := result.LastInsertId()
	catch.Id = int(lastId)
	go ipgeo.Process(catch.Id, catch.RemoteIp)
}

func CreateWithChan(catch Catch) <-chan error {
	c := make(chan error)

	go func(ec chan error) {
		_, err := mysql.Execute(`INSERT INTO CATCH(pot_id, remote_ip, create_at) VALUES (?, ?, ?)`, catch.PotId, catch.RemoteIp, catch.CreateDate)
		if err != nil {
			c <- err
		}

		c <- nil
	}(c)

	return nil
}
