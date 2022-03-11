package main

import (
	"fmt"
	"math"
)

func greet(name string) {
	fmt.Println("Hello there", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func cycleNames(names []string, f func(string)) {
	for i := 0; i < len(names); i++ {
		f(names[i])
	}
}

// Specifying return types
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	greet("Vettel")
	sayGoodbye("Alonso")

	names := []string{"Vettel", "Schumacher", "Alonso"}
	cycleNames(names, greet)
	cycleNames(names, sayGoodbye)

	area := circleArea(4.5)
	fmt.Println("Circle area", area)
}
