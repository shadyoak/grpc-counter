package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/shadyoak/grpc-counter/service"
	"google.golang.org/grpc"
)

type CounterServer struct {
	Counter Counter
	Port    int
}

func (s *CounterServer) IncrementCounter(stream service.Counter_IncrementCounterServer) error {
	for {

		in, err := stream.Recv()

		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Printf("receive error: %v", err)
			return err
		}

		count := s.Counter.Increment(in.Count)
		log.Println("current count:", count)

		val := &service.CounterValue{count}
		if err := stream.Send(val); err != nil {
			return err
		}
	}
}

func (c *CounterServer) Start() {
	port := fmt.Sprintf(":%v", c.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterCounterServer(s, &CounterServer{})
	log.Println("listening on port:", c.Port)
	s.Serve(lis)
}

func New(port int) *CounterServer {
	return &CounterServer{Port: port, Counter: 0}
}
