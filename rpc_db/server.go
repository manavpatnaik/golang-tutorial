package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type Args struct {
	A, B int
}

type API int
type Arithmetic int

var database []Item

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(item Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == item.Title {
			database[idx] = Item{item.Title, item.Body}
			changed = database[idx]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}

	*reply = del
	return nil
}

func (arith *Arithmetic) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	api := new(API)
	arith := new(Arithmetic)

	err := rpc.Register(api)
	rpc.Register(arith)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":4040", nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
