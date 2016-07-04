package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shadyoak/grpc-counter/client"
	"github.com/simon-whitehead/relayR"
)

type CountRelay struct {
}

type WebServer struct {
	exchange *relayr.Exchange
	rpcHost  string
	rpcPort  int
	webPort  int
}

func (w *WebServer) listenForUpdates(client *client.CounterClient) chan bool {
	c := make(chan bool)
	go func() {
	Loop:
		for {
			select {
			case count := <-client.Receive.CounterUpdate:
				w.exchange.Relay(CountRelay{}).Call("PushCount", count)
			case err := <-client.Receive.Error:
				log.Printf("rpc error: %v", err)
			case <-client.Receive.Done:
				log.Printf("updates complete")
				close(c)
				break Loop
			}
		}
	}()
	return c
}

func (tr CountRelay) PushCount(relay *relayr.Relay, count int) {
	relay.Clients.All("counterUpdated", count)
}

func New(rpcHost string, rpcPort, webServerPort int) *WebServer {
	exchange := relayr.NewExchange()
	exchange.RegisterRelay(CountRelay{})
	return &WebServer{
		exchange: exchange,
		rpcHost:  rpcHost,
		rpcPort:  rpcPort,
		webPort:  webServerPort,
	}
}

func (w *WebServer) Start() {

	// connect and listen to the RPC server
	client := client.New(w.rpcHost, w.rpcPort)
	if err := client.Start(); err != nil {
		log.Fatalf("unable to connect to rpc server: %v", err)
	}
	defer client.Stop()
	log.Println("web server connected to rpc server:", client.Address())
	w.listenForUpdates(client)

	// serve the website
	http.Handle("/relayr/", w.exchange)
	assets := assetFS()
	assets.Prefix += "/static/"

	// DEBUG: uncomment to serve from the file system directly
	//http.Handle("/", http.FileServer(http.Dir("./web/static")))
	http.Handle("/", http.FileServer(assets))

	addr := fmt.Sprintf(":%v", w.webPort)
	log.Println("web server listening on port:", w.webPort)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("unable to start web server: %v", err)
	}
}
