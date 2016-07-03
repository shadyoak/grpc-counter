package server

import (
	"log"
	"sync"

	"github.com/shadyoak/grpc-counter/service"
)

type counterClient interface {
	Send(m *service.CounterValue) error
}

type clientList struct {
	clients []counterClient
	mutex   sync.RWMutex
}

func (c *clientList) addClient(client counterClient) {
	c.mutex.Lock()
	c.clients = append(c.clients, client)
	c.mutex.Unlock()
	log.Println("client connected:", client)
}

func newClientList() clientList {
	c := []counterClient{}
	m := sync.RWMutex{}
	return clientList{c, m}
}

func (c *clientList) notifyAllClients(client counterClient, val service.CounterValue) error {
	c.mutex.RLock()
	// TODO: Needs a better error handling, concurrency, and timeouts...
	var result error
	for _, c := range c.clients {
		err := c.Send(&val)
		if err != nil && c == client {
			result = err
		}
	}
	c.mutex.RUnlock()
	return result
}

func (c *clientList) removeClient(client counterClient) {
	c.mutex.Lock()
	newClients := c.clients[:0]
	for _, c := range c.clients {
		if c != client {
			newClients = append(newClients, c)
		}
	}
	c.clients = newClients
	c.mutex.Unlock()
	log.Println("client disconnected:", client)
}
