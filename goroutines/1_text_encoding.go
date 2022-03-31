package main

import (
	"fmt"
	"strings"
)

func encode(input string, mychannel chan string) {
	var out string
	for _, val := range input {
		//fmt.Println(val,((val+1) % 97) + 97)
		out += string(((val + 1) % 97) + 97)
	}
	//fmt.Println(out)
	mychannel <- out
}

func hash(input string, encodedString string, mychannel chan int) {
	len := len(input)
	count := 0
	for _, val := range encodedString {
		//fmt.Println(val)
		if val == 97 || val == 101 || val == 105 || val == 111 || val == 117 {
			count = count + 1
		}
	}
	//fmt.Println(len)
	ans := (len) * count
	mychannel <- ans
}

func main() {
	var input string
	fmt.Println("Enter the text to be encoded")
	fmt.Scan(&input)
	mychannel1 := make(chan string)
	input = strings.ToLower(input)
	go encode(input, mychannel1)
	x := <-mychannel1
	mychannel2 := make(chan int)
	go hash(input, x, mychannel2)
	y := <-mychannel2
	fmt.Printf("%s_%d\n", x, y)
}
