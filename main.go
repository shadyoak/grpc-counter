package main

import "github.com/shadyoak/grpc-counter/server"

const (
	port = 5000
)

func main() {
	server := server.New(port)
	server.Start()
}
