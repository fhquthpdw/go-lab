package main

import "fmt"

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	// mac
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	// windows
	winComputer := &windows{}
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()
}
