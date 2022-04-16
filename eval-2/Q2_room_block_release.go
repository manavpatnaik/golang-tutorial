package main

import (
	"fmt"
)

func blockRoom(total int, n int, availableRooms chan int) {
	if n <= total {
		total = total - n
	}
	availableRooms <- total
}

func releaseRooms(currAvailable int, n int, rooms chan int) {
	rooms <- (currAvailable + n)
}

func main() {
	var totalRooms int
	fmt.Print("Enter the total number of rooms: ")
	fmt.Scan(&totalRooms)

	var blockCount int
	fmt.Print("Enter the number of rooms to be blocked: ")
	fmt.Scan(&blockCount)
	fmt.Println()

	availableRooms := make(chan int)

	currAvailable := totalRooms

	go blockRoom(totalRooms, blockCount, availableRooms)
	currAvailable = <-availableRooms
	fmt.Println("Blocked", blockCount, "rooms, Now available rooms:", currAvailable)
	fmt.Println()

	var releaseCount int
	fmt.Print("Enter the number of rooms to be released: ")
	fmt.Scan(&releaseCount)

	rooms := make(chan int)
	go releaseRooms(currAvailable, releaseCount, rooms)
	currAvailable = <-rooms
	if currAvailable <= totalRooms {
		fmt.Println("Released", releaseCount, "rooms, Now available rooms:", currAvailable)
	} else {
		fmt.Println("Cannot release those many rooms")
	}
}
