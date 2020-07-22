package main

import (
	"context"
	"fmt"
	"log"
	"net"

	user "github.com/jeffotoni/tdconline/grpc/proto"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Port string = ":8002"
	netl *net.Listener
)

type User struct{}

func (s *User) Get(ctx context.Context, in *user.Job) (*user.Reply, error) {
	fmt.Printf("\nIn Id:%v", string(in.Id))
	fmt.Printf("\nIn Name:%v", string(in.Name))
	res := &user.Reply{}
	res.Params = make(map[string][]byte)
	res.Params["jeffotoni"] = []byte("Meetup Campinas")
	return res, nil
}

func main() {
	netlis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("Error starting gRPC Worker Server:\n %s", err)
		return
	}

	netl = &netlis
	gRPCserver := ggrpc.NewServer()

	user.RegisterWorkerServiceServer(gRPCserver, &User{})
	log.Printf("gRPC Run %s", Port)
	reflection.Register(gRPCserver)
	if err := gRPCserver.Serve(*netl); err != nil {
		log.Fatalf("Error starting gRPC:\n %s", err)
	}
}
