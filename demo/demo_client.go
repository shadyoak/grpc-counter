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

	printChan := make(chan int)
	go printCounts(printChan)

	go func() {
	Loop:
		for {
			select {
			case count := <-client.Receive.CounterUpdate:
				printChan <- count
			case err := <-client.Receive.Error:
				close(printChan)
				grpclog.Fatalf("fatal error: %v", err)
			case <-client.Receive.Done:
				grpclog.Printf("updates complete")
				close(printChan)
				close(c)
				break Loop
			}
		}
	}()
	return c
}

func printCounts(c chan int) {

	flush := time.Tick(25 * time.Millisecond)

	var count *int
	for {
		select {
		case counter := <-c:
			count = &counter
		case <-flush:
			if count != nil {
				grpclog.Printf("got count: %v", *count)
				count = nil
			}
		}
	}
}

func main() {
	client := client.New(host, port)
	if err := client.Start(); err != nil {
		grpclog.Fatalf("unable to connect to client: %v", err)
	}
	defer client.Stop()
	grpclog.Println("connected to counter server:", client.Address())

	listen := listenForUpdates(client)

	for i := 0; i < 1000000; i++ {
		if err := client.IncrementCounter(1); err != nil {
			grpclog.Fatalf("increment counter error: %v", err)
		}
		//time.Sleep(500 * time.Millisecond)
	}

	<-listen
}
