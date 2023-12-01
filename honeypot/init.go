package honeypot

import (
	"github.com/tnuanchuay/honeypot/honeypot/catch"
	"github.com/tnuanchuay/honeypot/honeypot/ipgeo"
	"github.com/tnuanchuay/honeypot/honeypot/pot"
)

func Init() {
	pot.CreateTable()
	catch.CreateTable()
	ipgeo.CreateTable()
}
