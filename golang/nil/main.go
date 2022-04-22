package main

import "fmt"

func main() {
	ii := int(nil)
	ss := int(nil)

	if ss == ii {
		fmt.Println("ok")
	}
}
