package main

import "fmt"

func main() {
	adidasFactoryIns := getSportsFactory(AdidasFactory)
	nikeFactoryIns := getSportsFactory(NikeFactory)

	adidasShoe := adidasFactoryIns.makeShoe()
	adidasShirt := adidasFactoryIns.makeShirt()

	nikeShoe := nikeFactoryIns.makeShoe()
	nikeShirt := nikeFactoryIns.makeShirt()

	printShoe(adidasShoe)
	printShoe(nikeShoe)
	fmt.Println()
	printShirt(adidasShirt)
	printShirt(nikeShirt)
}

func printShoe(s iShoe) {
	fmt.Println("Logo: ", s.getShoeLogo())
	fmt.Println("Size: ", s.getShoeSize())
}
func printShirt(s iShirt) {
	fmt.Println("Logo: ", s.getShirtLogo())
	fmt.Println("Size: ", s.getShirtSize())
}
