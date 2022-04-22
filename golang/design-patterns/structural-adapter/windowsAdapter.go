package main

import "fmt"

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning single to USB.")
	w.windowMachine.insertIntoUSBPort()
}
