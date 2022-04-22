package main

import "fmt"

func main() {
	normalBuilderIns := getBuilder(NormalBuilderType)
	iglooBuilderIns := getBuilder(IglooBuilderType)

	// director build normal house
	directorNorIns := newDirector(normalBuilderIns)
	h1 := directorNorIns.buildHouse()
	fmt.Println("Normal House Door Type: ", h1.doorType)
	fmt.Println("Normal House Window Type: ", h1.windowType)
	fmt.Println("Normal House Floor: ", h1.floor)
	fmt.Println()

	// director build igloo house
	directorIloIns := newDirector(iglooBuilderIns)
	h2 := directorIloIns.buildHouse()
	fmt.Println("Igloo House Door Type: ", h2.doorType)
	fmt.Println("Igloo House Window Type: ", h2.windowType)
	fmt.Println("Igloo House Floor: ", h2.floor)
}
