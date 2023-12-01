package ipgeo

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

var client *IpGeoClient

func GetGeoLocation(ip string) (*IpGeoResult, error) {
	return client.GetGeoLocation(ip)
}

func Init() {
	client = New(viper.GetString("ipgeolocation.apiKey"))
}

type IpGeoClient struct {
	apiKey string
}

func New(apiKey string) *IpGeoClient {
	return &IpGeoClient{
		apiKey: apiKey,
	}
}

func (cli *IpGeoClient) GetGeoLocation(ip string) (*IpGeoResult, error) {
	res, err := http.Get(cli.getUrl(ip))
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var result IpGeoResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *IpGeoClient) getUrl(ip string) string {
	return fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", cli.apiKey, ip)
}
