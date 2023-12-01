package ipgeo

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/tnuanchuay/honeypot/ipgeo"
	"github.com/tnuanchuay/honeypot/mysql"
	"strconv"
	"time"
)

func Process(catchId int, remoteIp string) {
	result, err := ipgeo.GetGeoLocation(remoteIp)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug(remoteIp, result)

	lat, err := strconv.ParseFloat(result.Latitude, 64)
	if err != nil {
		lat = 0
		log.Error(err)
	}

	long, err := strconv.ParseFloat(result.Longitude, 64)
	if err != nil {
		long = 0
		log.Error(err)
	}

	ipGeoCatch := IpGeoCatch{
		CatchId:        catchId,
		CreateDate:     time.Now(),
		CountryName:    result.CountryName,
		StateProvince:  result.StateProv,
		ZipCode:        result.Zipcode,
		Latitude:       lat,
		Longitude:      long,
		Isp:            result.Isp,
		ConnectionType: result.ConnectionType,
		Organization:   result.Organization,
	}

	go Insert(ipGeoCatch)
}

func Insert(ipGeoCatch IpGeoCatch) {
	mysql.Execute(`
INSERT INTO GEOIP_CATCH (catch_id, create_at, country_name, state_prov, zipcode, latitude, longitude, isp, connection_type, organization)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`, ipGeoCatch.CatchId, ipGeoCatch.CreateDate, ipGeoCatch.CountryName, ipGeoCatch.StateProvince, ipGeoCatch.ZipCode, ipGeoCatch.Latitude, ipGeoCatch.Longitude, ipGeoCatch.Isp, ipGeoCatch.ConnectionType, ipGeoCatch.Organization)
}
