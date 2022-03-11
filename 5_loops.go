package main

import "fmt"

func main() {
	// A while loop
	x := 0
	for x < 5 {
		fmt.Print(x, " ")
		x++
	}
	fmt.Println()

	// The traditional for loop
	count := 5
	for i := 0; i < count; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Looping through an array/slice
	names := []string{"A", "B", "C", "D", "E", "F"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Looping using index and value
	for index, value := range names {
		fmt.Println(index, value)
	}

	// If we dont want to use the index, then use _
	for _, value := range names {
		fmt.Println(value)
	}

}
