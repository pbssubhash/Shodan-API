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

func (cred *Credential) Init() (*Credential, *APIInfo) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", cred.Url, cred.Key))
	body, err := ioutil.ReadAll(res.Body)
	var tesl *APIInfo
	json.Unmarshal([]byte(body), &tesl)
	if err != nil {
		return nil, nil
	}
	return cred, tesl
}

func Setup(apikey string, url string) (*Credential, *APIInfo) {
	creda := Credential{Url: url, Key: apikey}
	resp, keeda := creda.Init()
	return resp, keeda
}
