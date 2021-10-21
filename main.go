package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func main() {
	ips, err := net.LookupIP("iherb.okta.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("iherb.okta.com IN A %s\n", ip.String())
	}
	return

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://iherb.okta.com/oauth2/aus1iqzmrvfNssQO10h8/.well-known/openid-configuration", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(body)
}
