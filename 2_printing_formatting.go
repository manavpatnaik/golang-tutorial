package main

import "fmt"

func main() {
	// Print
	fmt.Print("Hello without newline at end")
	fmt.Print("Hello with newline\n")

	// Println
	fmt.Println("Hello people 1")
	fmt.Println("Hello people 2")

	age := 10
	name := "Scott"
	fmt.Println("My name is", name, "and I am", age)

	// Printf (formatted strings)
	fmt.Printf("My age: %d, My name: %s\n", age, name)
	fmt.Printf("My age: %q, My name: %q\n", age, name)
	fmt.Printf("My name: %q\n", name) // Adds quotes around string vars, doesnt work for numeric types
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)

	// Sprintf - save formatted strings, returns the formatted string
	str := fmt.Sprintf("My name: %s, my age: %d", name, age)
	fmt.Println(str)

	// More at: https://pkg.go.dev/fmt@go1.17.8
}
