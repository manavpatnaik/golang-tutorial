package main

import "fmt"

func updateName(s *string) {
	*s = "Updated"
}

func main() {
	name := "Manav"
	ptr := &name
	fmt.Println("Name:", name)
	fmt.Println("Name pointer:", ptr)
	fmt.Println("Name accessed using pointer:", *ptr)

	fmt.Println("Name before updating:", name)
	updateName(&name)
	fmt.Println("Updated name:", name)
}
