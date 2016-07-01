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
	ClientChan chan error
	Clients    []service.Counter_IncrementCounterServer
	Counter    Counter
	Port       int
}

func (s *CounterServer) addClient(stream service.Counter_IncrementCounterServer) {
	go func() {
		s.Clients = append(s.Clients, stream)
		s.ClientChan <- nil
	}()
	<-s.ClientChan

	log.Println("client connected:", stream)
}

func (s *CounterServer) removeClient(stream service.Counter_IncrementCounterServer) {
	go func() {
		clients := s.Clients
		newClients := clients[:0]
		for _, c := range clients {
			if c != stream {
				newClients = append(newClients, c)
			}
		}
		s.Clients = newClients
		s.ClientChan <- nil
	}()
	<-s.ClientChan

	log.Println("client disconnected:", stream)
}

func (s *CounterServer) notifyClients(stream service.Counter_IncrementCounterServer, val *service.CounterValue) error {

	go func() {
		var streamErr error
		for _, c := range s.Clients {
			err := c.Send(val)
			if err != nil && c == stream {
				streamErr = err
			}
		}
		s.ClientChan <- streamErr
	}()

	return <-s.ClientChan
}

func (s *CounterServer) IncrementCounter(stream service.Counter_IncrementCounterServer) error {

	s.addClient(stream)
	defer s.removeClient(stream)

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
		if err := s.notifyClients(stream, val); err != nil {
			return err
		}
	}
}

func (s *CounterServer) Start() {
	port := fmt.Sprintf(":%v", s.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service.RegisterCounterServer(grpcServer, s)
	log.Println("listening on port:", s.Port)
	grpcServer.Serve(lis)
}

func New(port int) *CounterServer {
	clientChan := make(chan error)
	clients := []service.Counter_IncrementCounterServer{}
	return &CounterServer{
		ClientChan: clientChan,
		Clients:    clients,
		Counter:    0,
		Port:       port}
}
