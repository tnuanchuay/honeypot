package ipgeo

type IpGeoResult struct {
	IP                  string `json:"ip"`
	ContinentCode       string `json:"continent_code"`
	ContinentName       string `json:"continent_name"`
	CountryCode2        string `json:"country_code2"`
	CountryCode3        string `json:"country_code3"`
	CountryName         string `json:"country_name"`
	CountryNameOfficial string `json:"country_name_official"`
	CountryCapital      string `json:"country_capital"`
	StateProv           string `json:"state_prov"`
	StateCode           string `json:"state_code"`
	District            string `json:"district"`
	City                string `json:"city"`
	Zipcode             string `json:"zipcode"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	IsEu                bool   `json:"is_eu"`
	CallingCode         string `json:"calling_code"`
	CountryTld          string `json:"country_tld"`
	Languages           string `json:"languages"`
	CountryFlag         string `json:"country_flag"`
	GeonameID           string `json:"geoname_id"`
	Isp                 string `json:"isp"`
	ConnectionType      string `json:"connection_type"`
	Organization        string `json:"organization"`
	Currency            struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	TimeZone struct {
		Name            string  `json:"name"`
		Offset          int     `json:"offset"`
		OffsetWithDst   int     `json:"offset_with_dst"`
		CurrentTime     string  `json:"current_time"`
		CurrentTimeUnix float64 `json:"current_time_unix"`
		IsDst           bool    `json:"is_dst"`
		DstSavings      int     `json:"dst_savings"`
	} `json:"time_zone"`
}
