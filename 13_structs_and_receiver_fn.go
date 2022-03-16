package main

import "fmt"

// A custom data structure
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Function to generate new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Receiver functions
// By doing this we limit this function only to bill objects
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("%-15v ...$%0.2f", "total:", total)
	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	// We dont need to explicitly dereference here too, GO does it by itself
	// b.tip = tip
	(*b).tip = tip
}

// Add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	myBill := newBill("Manav")
	fmt.Println(myBill.format())

	myBill.addItem("Item 1", 2.99)
	myBill.addItem("Item 2", 4.99)
	myBill.addItem("Item 3", 6.99)
	myBill.updateTip(20)
	fmt.Println(myBill.format())
}
