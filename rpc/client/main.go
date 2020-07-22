package main

import (
	"fmt"
	"log"
	"net/rpc"
)

var PORT_RPC = ":12346"

type Args struct {
	Json string
}

func main() {

	client, _ := rpc.Dial("tcp", PORT_RPC)

	// Synchronous call
	args := Args{`{"msg":"meu json para rpc tdc online!!"}`}
	var reply string

	// Calling my method
	err := client.Call("User.Get", args, &reply)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(reply)
}
