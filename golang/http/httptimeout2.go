package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := httpClient.Get("http://localhost:3030")
	if err != nil {
		if netErr, ok := err.(net.Error); ok {
			if netErr.Timeout() {
				fmt.Println("1")
				fmt.Println(err)
				fmt.Println()
			} else {
				fmt.Println("2")
				fmt.Println(netErr)
			}
		} else {
			fmt.Println("3")
			fmt.Println(err)
			fmt.Println()
		}
		return
	}
	fmt.Println("4")
	fmt.Println(response.StatusCode)
	a, err := ioutil.ReadAll(response.Body)
	fmt.Printf("%s", a)
}
