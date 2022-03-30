package main

import "fmt"

func main() {
	var opt int = 3
	switch opt {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number")
	}
}
