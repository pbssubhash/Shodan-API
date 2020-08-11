package Shodan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Credential struct {
	Url string
	Key string
}

type Response struct {
	Message string
	err     bool
}

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telner"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (cred *Credential) Init() (*Response, *APIInfo) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", cred.Url, cred.Key))
	body, err := ioutil.ReadAll(res.Body)
	var tesl APIInfo
	json.Unmarshal([]byte(body), &tesl)
	// if tesl.Unlocked == true {
	// 	fmt.Print("Unlocked")
	// }
	if err != nil {
		return &Response{"Failure", true}, nil
	}
	var respi *APIInfo
	err = json.NewDecoder(res.Body).Decode(&respi)
	return &Response{"Success", false}, respi
}

func Setup(apikey string, url string) (*Response, *Credential, *APIInfo) {
	// var creda *Credential
	creda := Credential{Url: url, Key: apikey}
	resp, keeda := creda.Init()
	return resp, &creda, keeda
}
