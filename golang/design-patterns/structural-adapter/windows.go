package main

import "fmt"

type windows struct {
}

func (m *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}
