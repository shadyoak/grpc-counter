package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/simon-whitehead/relayR"
)

type TimeRelay struct {
}

type WebServer struct {
	counterServerAddress string
	port                 int
}

func (tr TimeRelay) PushTime(relay *relayr.Relay) {
	relay.Clients.All("timeUpdated", time.Now().Local().Format("Mon Jan 2 2006 03:04:05 PM"))
}

func New(webServerPort int, counterServerAddress string) *WebServer {
	return &WebServer{
		counterServerAddress: counterServerAddress,
		port:                 webServerPort,
	}
}

func (w *WebServer) Start() {
	log.Println("web server listening on port:", w.port)
	exchange := relayr.NewExchange()
	exchange.RegisterRelay(TimeRelay{})

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				exchange.Relay(TimeRelay{}).Call("PushTime")
			}
		}
	}()

	http.Handle("/relayr/", exchange)

	assets := assetFS()
	assets.Prefix += "/static/"

	// DEBUG: uncomment to serve from the file system directly
	//http.Handle("/", http.FileServer(http.Dir("./web/static")))
	http.Handle("/", http.FileServer(assets))

	addr := fmt.Sprintf(":%v", w.port)
	log.Println("web server listening at address:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("unable to start web server: %v", err)
	}
}
