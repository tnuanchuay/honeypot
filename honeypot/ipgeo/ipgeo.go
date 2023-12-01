package ipgeo

import (
	"github.com/tnuanchuay/honeypot/mysql"
	"time"
)

type IpGeoCatch struct {
	Id             int       `json:"id" form:"id" query:"id"`
	CatchId        int       `json:"catch_id" form:"catch_id" query:"catch_id"`
	CreateDate     time.Time `json:"create_at" form:"create_date" query:"create_date"`
	CountryName    string    `json:"country_name" form:"country_name" query:"country_name"`
	StateProvince  string    `json:"state_prov" form:"state_prov" query:"state_prov"`
	ZipCode        string    `json:"zipcode" form:"zipcode" query:"zipcode"`
	Latitude       float64   `json:"latitude" form:"latitude" query:"latitude"`
	Longitude      float64   `json:"longitude" form:"longitude" query:"longitude"`
	Isp            string    `json:"isp" form:"isp" query:"isp"`
	ConnectionType string    `json:"connection_type" form:"connection_type" query:"connection_type"`
	Organization   string    `json:"organization" form:"organization" query:"organization"`
}

func CreateTable() {
	_, err := mysql.Execute(`
CREATE TABLE IF NOT EXISTS GEOIP_CATCH (
    id	INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    catch_id INT NOT NULL,
    create_at	TIMESTAMP		NOT NULL,
    country_name VARCHAR(100),
    state_prov VARCHAR(100),
    zipcode VARCHAR(10),
    latitude DOUBLE,
    longitude DOUBLE,
    isp VARCHAR(500),
    connection_type VARCHAR(100),
    organization VARCHAR(500)
    )
`)
	if err != nil {
		panic(err)
	}
}
