package main

import "fmt"

func main() {
	age := 45

	fmt.Println("age == 50", age == 50)
	fmt.Println("age != 50", age != 50)
	fmt.Println("age < 50", age < 50)
	fmt.Println("age <= 50", age <= 50)
	fmt.Println("age > 50", age > 50)
	fmt.Println("age >= 50", age >= 50)

	if age == 45 {
		fmt.Println("Age is 45")
	} else if age < 45 {
		fmt.Println("Age < 45")
	} else {
		fmt.Println("Age > 45")
	}

	names := []string{"A", "B", "C", "D", "E", "F"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		} else if index == 4 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
