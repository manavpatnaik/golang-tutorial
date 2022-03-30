package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   3.99,
		"salad": 1.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["soup"])
	fmt.Println(menu["pie"])
	fmt.Println(menu["salad"])

	// Loop through maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Another map
	phonebook := map[int]string{
		1234567: "Manav",
		7654321: "John",
		7652345: "Patnaik",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567])

	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}

	// Update map items
	phonebook[1234567] = "Manav Patnaik"
	fmt.Println(phonebook[1234567])

	delete(phonebook, 1234567)
	fmt.Println(phonebook)

}
