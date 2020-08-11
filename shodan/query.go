package Shodan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	City        string `json:"city"`
	RegionCode  string `json:"region_code"`
	AreaCode    string `json:"area_code"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	CountryCode int    `json:"country_code3"`
	CountryName string `json:"country_name"`
	PostalCode  int    `json:"postal_code"`
}
type Host struct {
	OS           string   `json:"os"`
	ISP          string   `json:"isp"`
	hostname     []string `json:"hostnames"`
	location     Location `json:"location"`
	ip           int      `json:"ip_str"`
	domains      []string `json:"domains"`
	Port         int      `json:"port"`
	Data         string   `json:"data"`
	Organisation string   `json:"org"`
}

type Matches struct {
	Matches []Host `json:"matches"`
}

func Important() {
	fmt.Println("Important")
}

func QueryShodan(query string, cred Credential) *Matches {
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", cred.Url, cred.Key, query))
	if err != nil {
		log.Fatal("Error-3")
	}
	var tesl *Matches
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tesl)
	if err != nil {
		log.Fatal("Error")
	}
	return tesl
}
