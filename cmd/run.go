package cmd

import (
	"github.com/tnuanchuay/honeypot/app"
	"github.com/tnuanchuay/honeypot/config"
	"github.com/tnuanchuay/honeypot/honeypot"
	"github.com/tnuanchuay/honeypot/log"
	"github.com/tnuanchuay/honeypot/mysql"
)

func Run() {
	log.Init()
	config.Init()
	mysql.InitWithDefault()
	honeypot.Init()
	app.Init()
	app.Get("/honey/new", honeypot.CreateGetHandler())
	app.Run()
}
