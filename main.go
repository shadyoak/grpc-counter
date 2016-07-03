package main

import "github.com/shadyoak/grpc-counter/server"

const (
	defaultPort = 5000
)

func main() {
	server := server.New(defaultPort)
	server.Start()
}
