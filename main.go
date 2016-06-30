package main

import (
	"log"
	"net"

	"github.com/shadyoak/grpc-counter/service"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

type counterServer struct{}

func (s *counterServer) IncrementCounter(stream service.Counter_IncrementCounterServer) error {
	// for {
	// 	in, err := stream.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// 	key := serialize(in.Location)
	// 	if _, present := s.routeNotes[key]; !present {
	// 		s.routeNotes[key] = []*pb.RouteNote{in}
	// 	} else {
	// 		s.routeNotes[key] = append(s.routeNotes[key], in)
	// 	}
	// 	for _, note := range s.routeNotes[key] {
	// 		if err := stream.Send(note); err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	c := service.CounterValue{2}
	log.Println(c)

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterCounterServer(s, &counterServer{})
	log.Println("Listenting on port", port)
	s.Serve(lis)
}
