package main

import (
	"time"

	"github.com/shadyoak/grpc-counter/client"
	"google.golang.org/grpc/grpclog"
)

const (
	host = "localhost"
	port = 5000
)

func listenForUpdates(client *client.CounterClient) chan bool {
	c := make(chan bool)
	go func() {
	Loop:
		for {
			select {
			case count := <-client.Receive.CounterUpdate:
				grpclog.Printf("got count: %v", count)
			case err := <-client.Receive.Error:
				grpclog.Fatalf("fatal error: %v", err)
			case <-client.Receive.Done:
				grpclog.Printf("updates complete")
				close(c)
				break Loop
			}
		}
	}()
	return c
}

func main() {
	client := client.New(host, port)
	err := client.Start()
	if err != nil {
		grpclog.Fatalf("unable to connect to client: %v", err)
	}
	defer client.Stop()
	grpclog.Println("connected to counter server:", client.Address())

	listen := listenForUpdates(client)

	for i := 0; i < 10; i++ {
		if err := client.IncrementCounter(1); err != nil {
			grpclog.Fatalf("increment counter error: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	<-listen
}
