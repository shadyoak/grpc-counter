package main

import (
	"github.com/shadyoak/grpc-counter/client"
	"google.golang.org/grpc/grpclog"
)

const (
	host = "localhost"
	port = 5000
)

func main() {
	client := client.New(host, port)
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
