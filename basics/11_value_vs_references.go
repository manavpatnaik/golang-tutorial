package main

import "fmt"

// GO is a "pass-by-value" language
func updateName(s string) {
	s = "Updated"
}

// Returning values
func updateNameV1(s string) string {
	s = "Updated"
	return s
}

// Updating non primitive data types
func updateMap(m map[string]float64) {
	m["tea"] = 10
}

func main() {
	name := "Manav"

	fmt.Println(name)
	updateName(name)
	// 'name' doesnt get updated
	fmt.Println(name)

	name = updateNameV1(name)
	// Now its updated
	fmt.Println(name)

	menu := map[string]float64{
		"tea": 20,
	}
	fmt.Println(menu)
	updateMap(menu)
	// Now the map is updated as its pass by reference
	fmt.Println(menu)
}
