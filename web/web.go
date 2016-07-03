package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/simon-whitehead/relayR"
)

type CountRelay struct {
}

type WebServer struct {
	counterServerAddress string
	exchange             *relayr.Exchange
	port                 int
}

func (tr CountRelay) PushCount(relay *relayr.Relay, count int) {
	relay.Clients.All("counterUpdated", count)
}

func New(webServerPort int, counterServerAddress string) *WebServer {
	exchange := relayr.NewExchange()
	exchange.RegisterRelay(CountRelay{})
	return &WebServer{
		counterServerAddress: counterServerAddress,
		exchange:             exchange,
		port:                 webServerPort,
	}
}

func (w *WebServer) Start() {
	log.Println("web server listening on port:", w.port)

	counter := 0

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				counter += 1
				w.exchange.Relay(CountRelay{}).Call("PushCount", counter)
			}
		}
	}()

	http.Handle("/relayr/", w.exchange)
	assets := assetFS()
	assets.Prefix += "/static/"

	// DEBUG: uncomment to serve from the file system directly
	//http.Handle("/", http.FileServer(http.Dir("./web/static")))
	http.Handle("/", http.FileServer(assets))

	addr := fmt.Sprintf(":%v", w.port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("unable to start web server: %v", err)
	}
}
