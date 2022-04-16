package main

import (
	"log"
	"net/http"
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

type Supermarket int

var db []Item
var balance int = 100

func (s *Supermarket) GetDB(empty string, reply *[]Item) error {
	*reply = db
	return nil
}

func (s *Supermarket) OrderItem(order Order, reply *OrderStatus) error {
	var status OrderStatus

	for _, val := range db {
		if val.Item_name == order.Item_name {
			if val.Qty >= order.Qty {
				status = OrderStatus{"ORDER_ACCEPTED", order.Qty * val.Price}
			} else {
				status = OrderStatus{"REJECTED", 0}
			}
		}
	}
	*reply = status
	return nil
}

func (s *Supermarket) SellItem(item Item, reply *string) error {
	status := "OK"
	if item.Price*item.Qty <= balance {
		status = "OK"
	} else {
		status = "Cannot sell Item"
	}
	*reply = status
	return nil
}

func (s *Supermarket) AddItem(item Item, reply *Item) error {
	db = append(db, item)
	*reply = item
	return nil
}

func main() {
	supermarket := new(Supermarket)

	err := rpc.Register(supermarket)

	if err != nil {
		log.Fatal("error registering Supermarket", err)
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":4040", nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
