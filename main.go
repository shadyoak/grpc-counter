package main

import "github.com/shadyoak/grpc-counter/server"

func main() {
	server := server.New()
	server.Start()
}
