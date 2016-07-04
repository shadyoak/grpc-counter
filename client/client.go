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
	cleanup  func()
	hostname string
	port     int
	Receive  ReceiveChannels
}

type ReceiveChannels struct {
	CounterUpdate chan int
	Error         chan error
	Done          chan bool
}

func (c *CounterClient) countToTen(stream service.Counter_IncrementCounterClient) <-chan bool {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			counter := &service.CounterValue{int32(1)}
			if err := stream.Send(counter); err != nil {
				grpclog.Fatalf("failed to send a count: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}

func (c *CounterClient) listenForUpdates(stream service.Counter_IncrementCounterClient) <-chan bool {
	ch := make(chan bool)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				c.Receive.Done <- true
				return
			} else if err != nil {
				c.Receive.Error <- err
			} else {
				count := int(in.Count)
				c.Receive.CounterUpdate <- count
			}
		}
	}()
	return ch
}

func (c *CounterClient) beginStreaming(client service.CounterClient) service.Counter_IncrementCounterClient {

	stream, err := client.IncrementCounter(context.Background())
	if err != nil {
		grpclog.Fatalf("unable to create stream: %v", err)
	}

	c.listenForUpdates(stream)
	c.countToTen(stream)

	return stream
}

func New(hostname string, port int) *CounterClient {
	receive := ReceiveChannels{
		CounterUpdate: make(chan int),
		Error:         make(chan error),
		Done:          make(chan bool),
	}
	return &CounterClient{
		hostname: hostname,
		port:     port,
		Receive:  receive,
	}
}

func (c *CounterClient) Start() {
	address := fmt.Sprintf("%v:%v", c.hostname, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("unable to connect: %v", err)
	}
	grpclog.Println("connected to:", address)

	client := service.NewCounterClient(conn)
	stream := c.beginStreaming(client)

	c.cleanup = func() {
		stream.CloseSend()
		conn.Close()
	}
}

func (c *CounterClient) Stop() {
	c.cleanup()
}

func main() {
	client := New(defaultHostname, defaultPort)
	client.Start()
	defer client.Stop()

Loop:
	for {
		select {
		case count := <-client.Receive.CounterUpdate:
			grpclog.Printf("got count: %v", count)
		case err := <-client.Receive.Error:
			grpclog.Fatalf("fatal error: %v", err)
		case <-client.Receive.Done:
			grpclog.Printf("updates complete")
			break Loop
		}
	}
}
