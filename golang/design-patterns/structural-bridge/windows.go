package main

import "fmt"

type windows struct {
	printer printer
}

func (m *windows) print() {
	fmt.Println("Print request for windows")
	m.printer.printFile()
}

func (m *windows) setPrinter(p printer) {
	m.printer = p
}
