package main

import (
	"flag"
	"fmt"

	Shodan "github.com/pbssubhash/Shodan-API/shodan"
)

func main() {
	fmt.Println("[+] Welcome to Shodan GO Client [-]")
	fmt.Println("[-] Built by zer0 p1k4chu for learning purposes. [-]")
	const url = "https://api.shodan.io"
	key := flag.String("k", "NotPresent", "API Key for authenticating API requests.")
	help := flag.Bool("h", false, "Print Help")
	flag.Parse()
	if *key == "NotPresent" || *help == true {
		flag.PrintDefaults()
		return
	}
	resp, _, keeda := Shodan.Setup(*key, url)
	fmt.Println(resp.Message, keeda.Unlocked)
}
