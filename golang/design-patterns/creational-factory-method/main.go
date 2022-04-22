package main

import "fmt"

func main() {
	ak47 := getGun(GunTypeAk47)
	musket := getGun(GunTypeMusket)

	printDetail(ak47)
	fmt.Println()
	printDetail(musket)
}

func printDetail(g iGun) {
	fmt.Println("Gun: ", g.getName())
	fmt.Println("Power: ", g.getPower())
}
