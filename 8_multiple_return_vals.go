package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	return initials[0], initials[1]
}

func main() {
	fmt.Println(getInitials("Manav Patnaik"))
	fmt.Println(getInitials("Gaurav Patnaik"))

	// Declared in another file, but in the same package
	// Run using: go run 8_multiple_return_vals.go 9_package_scope.go
	fmt.Println(points)
	sayHello("Manav")
}
