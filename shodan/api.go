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

func (cred *Credential) Init() *Response {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", cred.Url, cred.Key))
	body, err := ioutil.ReadAll(res.Body)
	var tesl APIInfo
	json.Unmarshal([]byte(body), &tesl)
	// if tesl.Unlocked == true {
	// 	fmt.Print("Unlocked")
	// }
	if err != nil {
		return &(Response{"Failure", true})
	}
	return &(Response{"Success", false})
}

func Setup(apikey string, url string) *Response {
	var Creda Credential
	Creda = Credential{Url: url, Key: apikey}
	resp := Creda.Init()
	return resp

}