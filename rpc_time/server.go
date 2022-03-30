package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)

	listener, err := net.Listen("tcp", "0.0.0.0:1234")

	if err != nil {
		log.Fatal("Listener error: ", err)
	}

	http.Serve(listener, nil)
}

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Set the value at the pointer got from the client
	*reply = time.Now().Unix()
	return nil
}
