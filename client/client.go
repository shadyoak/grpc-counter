package main

import (
	"io"
	"log"
	"time"

	"github.com/shadyoak/grpc-counter/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	address = "localhost:5001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()
	client := service.NewCounterClient(conn)
	testCounter(client)
}

func testCounter(client service.CounterClient) {

	stream, err := client.IncrementCounter(context.Background())
	if err != nil {
		grpclog.Fatalf("%v.IncrementCounter(_) = _, %v", client, err)
	}

	waitc := make(chan struct{})

	go func(c chan struct{}) {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(c)
				return
			}
			if err != nil {
				grpclog.Fatalf("failed to receive count: %v", err)
			}
			grpclog.Printf("got count: %v", in.Count)
		}
	}(waitc)

	for i := 0; i < 10; i++ {
		c := &service.CounterValue{int32(i)}
		if err := stream.Send(c); err != nil {
			grpclog.Fatalf("failed to send a count: %v", err)
		}

		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-waitc
}
