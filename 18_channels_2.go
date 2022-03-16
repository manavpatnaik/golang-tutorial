package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Msg 1 from goroutine: going to sleep for 2 seconds")
	time.Sleep(4 * time.Second)
	fmt.Println("Msg 2 from goroutine: goroutine awake, task done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main is going to call goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}
