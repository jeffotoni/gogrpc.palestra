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
	args := Args{`{"msg":"meu json para rpc Meetup Campinas!","name":"Marcos","age":22}`}
	var reply string
	for i:=0; i<=1000;i++ {
		// Calling my method
		err := client.Call("User.Get", args, &reply)
		if err != nil {
			log.Println(err)
			continue
		}
		println(".......")
	}
	fmt.Println(reply)
}
