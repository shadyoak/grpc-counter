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
	clients clientList
	counter counter
	port    int
}

func (s *CounterServer) IncrementCounter(stream service.Counter_IncrementCounterServer) error {
	s.clients.addClient(stream)
	defer s.clients.removeClient(stream)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Printf("receive error: %v", err)
			return err
		}

		count := s.counter.increment(in.Count)
		log.Println("current count:", count)
		val := service.CounterValue{count}

		if err := s.clients.notifyAllClients(stream, val); err != nil {
			log.Println("notify error:", err)
			return err
		}
	}
}

func New(port int) *CounterServer {
	return &CounterServer{
		clients: newClientList(),
		counter: 0,
		port:    port,
	}
}

func (s *CounterServer) Start() {
	port := fmt.Sprintf(":%v", s.port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service.RegisterCounterServer(grpcServer, s)
	log.Println("listening on port:", s.port)
	grpcServer.Serve(lis)
}
