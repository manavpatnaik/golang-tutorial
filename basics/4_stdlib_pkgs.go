package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello There!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.Contains(greeting, "hello"))

	// No in-place changes, it returns a new one
	updatedString := strings.ReplaceAll(greeting, "Hello", "Hey")
	fmt.Println(greeting, updatedString)

	// Some other functions
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ToLower(greeting))

	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Index(greeting, "dfg"))

	fmt.Println(strings.Split(greeting, ""))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(strings.Split(greeting, "e"))

	// Performs in-place sorting
	ages := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 4)
	fmt.Println(index)

	names := []string{"z", "y", "x", "w", "v", "u"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "u"))
}
