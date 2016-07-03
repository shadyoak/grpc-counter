package web

import (
	"fmt"
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
	http.Handle("/", http.FileServer(http.Dir("./web")))

	addr := fmt.Sprintf(":%v", w.port)
	http.ListenAndServe(addr, nil)
}

func main() {
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
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":80", nil)
}
