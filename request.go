package iplocate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IpLocate struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	AsName        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

func Request(ip string) (*IpLocate, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result IpLocate
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("ip-api error for %s", ip)
	}

	return &result, nil
}