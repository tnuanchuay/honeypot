package cmd

import (
	"github.com/tnuanchuay/honeypot/app"
	"github.com/tnuanchuay/honeypot/config"
	"github.com/tnuanchuay/honeypot/honeypot"
	"github.com/tnuanchuay/honeypot/mysql"
)

func Run() {
	config.Init()
	mysql.InitWithConfig()
	honeypot.Init()
	app.Init()
	app.Run()
}
