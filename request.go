package IpLocate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpLocate struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countrycode"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	Region        string  `json:"region"`
	District      string  `json:"district"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	ZipCode       string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	TimeZone      string  `json:"timezone"`
	OffSet        int     `json:"offset"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	AsName        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

var data IpLocate

func Request(ip string) {
	url := "http://ip-api.com/json/" + ip

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &data)
	
}
func GetQuery() string {
	return data.Query
}
func GetContinent() string {
	return data.Continent
}
func GetContinentCode() string {
	return data.ContinentCode
}
func GetCountry() string {
	return data.Country
}
func GetCountryCode() string {
	return data.CountryCode
}
func GetRegion() string {
	return data.Region
}
func GetRegionName() string {
	return data.RegionName
}
func GetDistrict() string {
	return data.District
}
func GetCity() string {
	return data.City
}
func GetZipCode() string {
	return data.ZipCode
}
func GetLat() float64 {
	return data.Lat
}
func GetLon() float64 {
	return data.Lon
}
func GetTimeZone() string {
	return data.TimeZone
}
func GetOffset() int {
	return data.OffSet
}
func GetHosting() bool {
	return data.Hosting
}
func GetProxy() bool {
	return data.Proxy
}
func GetMobile() bool {
	return data.Mobile
}
func GetIsp() string {
	return data.Isp
}
func GetOrg() string {
	return data.Org
}
func GetAs() string {
	return data.As
}
