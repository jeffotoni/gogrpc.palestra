package main

import (
	"context"
	"fmt"
	"time"

	gproto "github.com/jeffotoni/tdconline/grpc/proto"
	ggrpc "google.golang.org/grpc"
)

type Client struct {
	Host    string
	Timeout time.Duration
}

func (c *Client) Work(job *gproto.Job) (res *gproto.Reply, err error) {
	conn, err := ggrpc.Dial(c.Host, ggrpc.WithInsecure())
	fmt.Printf("\nNew connection %v, timeout[%s]", c.Host, c.Timeout)
	if err != nil {
		fmt.Printf("Error dial gRPC Worker client \n error[%v]", err)
		return nil, err
	}
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	w := gproto.NewWorkerServiceClient(conn)
	res, err = w.Get(ctx, job)
	if err != nil {
		fmt.Printf("Error performing gRPC work error[%v] \n job[%v]", err, job)
		return res, err
	}
	return res, nil
}

func main() {
	var c = Client{"localhost:8002", time.Second * 30}
	fmt.Println(c.Work(createJob()))
}

func createJob() *gproto.Job {
	job := &engdb_pb.Job{}
	job.Name = "jeffotoni..2020..."
	job.Id = "123469696934696xx"
	return job
}
