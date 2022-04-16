package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Item_name string
	Qty       int
	Price     int
}

type Order struct {
	Item_name string
	Qty       int
}

type OrderStatus struct {
	Status string
	Amount int
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client.Call("Supermarket.AddItem", Item{"A", 10, 15}, &reply)
	client.Call("Supermarket.AddItem", Item{"B", 20, 20}, &reply)
	client.Call("Supermarket.AddItem", Item{"C", 30, 25}, &reply)
	client.Call("Supermarket.AddItem", Item{"D", 40, 30}, &reply)
	client.Call("Supermarket.GetDB", "empty", &db)

	fmt.Println("Database: ", db)
	fmt.Println()

	var status OrderStatus
	ord1 := Order{"A", 5}
	ord2 := Order{"B", 10}
	ord3 := Order{"C", 60}
	client.Call("Supermarket.OrderItem", ord1, &status)
	fmt.Println("Order:", ord1, "Status:", status)
	client.Call("Supermarket.OrderItem", ord2, &status)
	fmt.Println("Order:", ord2, "Status:", status)
	client.Call("Supermarket.OrderItem", ord3, &status)
	fmt.Println("Order:", ord3, "Status:", status)

	fmt.Println()

	var saleStatus string
	item1 := Item{"X", 5, 10}
	client.Call("Supermarket.SellItem", item1, &saleStatus)
	fmt.Println("Item to be sold:", item1, saleStatus)
	item2 := Item{"Y", 12, 10}
	client.Call("Supermarket.SellItem", item2, &saleStatus)
	fmt.Println("Item to be sold:", item2, saleStatus)
}
