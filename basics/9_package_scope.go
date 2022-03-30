package main

import "fmt"

// These are accessible in any .go file which has package as main

var points = []int{10, 20, 30, 40, 50}

func sayHello(n string) {
	fmt.Println("Hello", n)
}
