package main

import (
	"fmt"
	"io"
	"time"

	"github.com/shadyoak/grpc-counter/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	defaultHostname = "localhost"
	defaultPort     = 5000
)

type CounterClient struct {
	hostname string
	port     int
}

func countToTen(stream service.Counter_IncrementCounterClient) <-chan bool {
	c := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c := &service.CounterValue{int32(1)}
			if err := stream.Send(c); err != nil {
				grpclog.Fatalf("failed to send a count: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return c
}

func listenForUpdates(stream service.Counter_IncrementCounterClient) <-chan bool {
	c := make(chan bool)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(c)
				return
			} else if err != nil {
				grpclog.Fatalf("failed to receive count: %v", err)
			}
			grpclog.Printf("got count: %v", in.Count)
		}
	}()
	return c
}

func runTest(client service.CounterClient) {

	stream, err := client.IncrementCounter(context.Background())
	if err != nil {
		grpclog.Fatalf("unable to create stream: %v", err)
	}
	defer stream.CloseSend()

	listen := listenForUpdates(stream)
	count := countToTen(stream)

	<-count
	<-listen
}

func New(hostname string, port int) *CounterClient {
	return &CounterClient{
		hostname: hostname,
		port:     port,
	}
}

func (c *CounterClient) Start() {
	address := fmt.Sprintf("%v:%v", c.hostname, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("unable to connect: %v", err)
	}
	grpclog.Println("connected to:", address)
	defer conn.Close()
	client := service.NewCounterClient(conn)
	runTest(client)
}

func main() {
	client := New(defaultHostname, defaultPort)
	client.Start()
}
