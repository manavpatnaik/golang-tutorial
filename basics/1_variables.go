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

	// Variations of the int - bits and memory
	var numOne int8 = 25 // Range: [-2^7, 2^7-1] i.e [-128, 127]
	var numTwo int16 = 26
	var numThree = 129
	var numFour uint = 12
	var numFive uint8 = 13
	fmt.Println(numOne, numTwo, numThree)
	fmt.Println(numFour, numFive)

	// Floats
	var scoreOne float32 = 24
	var scoreTwo float64 = 24234324 // Better to use float64, higher precision
	scoreThree := 3.4               // By default floats will be inferred at float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Reference: https://go.dev/ref/spec#Numeric_types
}
