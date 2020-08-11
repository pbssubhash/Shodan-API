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
	query := flag.String("q", "Default", "Query Shodan. Eg: product:apache")
	help := flag.Bool("h", false, "Print Help")
	flag.Parse()
	if *key == "NotPresent" || *help == true || *query == "Default" {
		flag.PrintDefaults()
		return
	}
	cred, _ := Shodan.Setup(*key, url)
	Shodan.QueryShodan(*query, *cred)
}
