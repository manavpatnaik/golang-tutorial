package main

import "fmt"

func main() {
	var ages [3]int = [3]int{1, 2, 3}
	fmt.Println("Hello", ages)

	// Automatic inference
	var nums = [3]int{1, 2, 3}
	fmt.Println(nums)

	// Array methods
	fmt.Println("Ages:", ages)
	fmt.Println("Length of Ages:", len(ages))

	// String array
	names := [4]string{"A", "B", "C", "D"}
	fmt.Println(names, len(names))
	names[2] = "b"
	fmt.Println(names)

	// Slices, appending is possible
	var scores = []int{10, 20, 30, 40, 50}
	fmt.Println(scores, len(scores))
	// append(slice, element) => returns a new slice
	scores = append(scores, 60)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Ccc")
	fmt.Println(rangeOne)
}
