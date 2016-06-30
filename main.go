package main

import (
	"io"
	"log"
	"net"

	"github.com/shadyoak/grpc-counter/service"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

type CounterServer struct {
	CurrentCount int
}

func (s *CounterServer) IncrementCounter(stream service.Counter_IncrementCounterServer) error {
	for {

		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		key := in.Count
		log.Println(s.CurrentCount, key)

		if err := stream.Send(&service.CounterValue{key}); err != nil {
			return err
		}

	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterCounterServer(s, &CounterServer{})
	log.Println("Listening on port", port)
	s.Serve(lis)
}
