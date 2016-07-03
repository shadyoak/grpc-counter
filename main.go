package main

import (
	"github.com/shadyoak/grpc-counter/server"
	"github.com/shadyoak/grpc-counter/web"
)

const (
	rpcHost = "localhost"
	rpcPort = 5000
	webPort = 8080
)

func main() {

	wait := make(chan bool)

	// start the counter RPC server
	go func() {
		server := server.New(rpcPort)
		server.Start()
		close(wait)
	}()

	// start the web server
	go func() {
		web := web.New(rpcHost, rpcPort, webPort)
		web.Start()
		close(wait)
	}()

	<-wait
}
