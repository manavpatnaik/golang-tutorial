package main

import "fmt"

func main() {
	// Strings
	var nameOne string = "Manav"
	var nameTwo = "Patnaik"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "New string 1"
	nameTwo = "New string 2"
	nameThree = "New string 3"
	fmt.Println(nameOne, nameTwo, nameThree)

	// Another way of declaring variables. ':' replaces the 'var'
	nameFour := "manav"
	fmt.Println(nameFour)

	// Ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)
}
