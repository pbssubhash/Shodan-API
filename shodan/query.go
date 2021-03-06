package Shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	City        string `json:"city,omitempty"`
	RegionCode  string `json:"region_code,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	CountryCode int    `json:"country_code3,omitempty"`
	CountryName string `json:"country_name,omitempty"`
	PostalCode  int    `json:"postal_code,omitempty"`
}
type Host struct {
	OS           string   `json:"os,omitempty"`
	ISP          string   `json:"isp,omitempty"`
	Hostname     []string `json:"hostnames,omitempty"`
	Location     Location `json:"location,omitempty"`
	IP           string   `json:"ip_str,omitempty"`
	Domains      []string `json:"domains,omitempty"`
	Port         int      `json:"port,omitempty"`
	Organisation string   `json:"org,omitempty"`
}

type Matches struct {
	Matches []Host `json:"matches"`
}

func Important() {
	fmt.Println("Important")
}

func QueryShodan(query string, cred Credential) *Matches {
	fmt.Println(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", cred.Url, cred.Key, query))
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", cred.Url, cred.Key, query))
	if err != nil {
		log.Fatal("Error-3")
	}
	var tesm Matches
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp.Status)
	// if err != nil {
	// 	log.Fatal("Error-4")
	// }
	json.NewDecoder(resp.Body).Decode(&tesm)
	// if err != nil {
	// 	log.Fatal("Error")
	// }
	return &tesm
}
