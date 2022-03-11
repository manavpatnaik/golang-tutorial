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
}
