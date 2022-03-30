package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello World goroutine")
	done <- true
}

func main() {
	// Conventional method of declaring
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	// Shorthand method
	// b := make(chan int)

	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("Main function")
}
