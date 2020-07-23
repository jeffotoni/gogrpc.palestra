// Go in Action
// @jeffotoni

package main

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"time"
)

var PORT_RPC = ":12346"

type Args struct {
	Json string
}

type User struct{}

/* {
	"msg":"meu json para rpc Meetup Campinas!",
	"name":Marcos","age":"22"
	}
 */
type ClientS struct {
	Msg string
	Name string
	Age int
}

func (t *User) Get(args *Args, reply *string) error {
	if len(args.Json) <= 0 {
		*reply = `{"status":"error", "msg":"json field is required"}`
		return nil
	}
	var c ClientS
	err := json.Unmarshal([]byte(args.Json), &c)
	if err!=nil {
		log.Println("Error Unmarshal:", err)
		return err
	}
	println(".......................................")
	println("::::::::json:::", args.Json)
	println("::::::::Name:::", c.Msg)
	println("::::::::Name:::", c.Name)
	println("::::::::Name:::", c.Age)
	println(".......................................")
	// ... faz algo aqui..
	*reply = `{"status":"ok", "msg":"User.Get json","id":"38xxx9w9x.."}`
	return nil
}

func Rpc() {
	re := new(User)
	serverRpc := rpc.NewServer()
	serverRpc.Register(re)
	tcpAddr, err := net.ResolveTCPAddr("tcp", PORT_RPC)
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	log.Println("Run RPC port:", PORT_RPC)
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Println(err.Error())
			time.Sleep(time.Second * 5)
			continue
		} else {
			serverRpc.ServeConn(conn)
		}
	}
}

func main() {
	Rpc()
}
