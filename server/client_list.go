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

func (c *clientList) notifyAllClients(client counterClient, count int) error {

	clients := make(chan counterClient)
	errors := make(chan error)
	errorList := make([]error, 0)

	wg := sync.WaitGroup{}
	wg.Add(4)
	go sendCount(count, clients, errors, &wg)
	go sendCount(count, clients, errors, &wg)
	go sendCount(count, clients, errors, &wg)
	go sendCount(count, clients, errors, &wg)
	go processErrors(errors, &errorList)

	c.mutex.RLock()
	for _, c := range c.clients {
		clients <- c
	}
	close(clients)
	c.mutex.RUnlock()

	wg.Wait()
	close(errors)

	if len(errorList) > 0 {
		return errorList[0]
	}

	return nil
}

func processErrors(errors chan error, acc *[]error) {
Loop:
	for {
		select {
		case err, ok := <-errors:
			if ok {
				*acc = append(*acc, err)
			} else {
				break Loop
			}
		}
	}
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

func sendCount(count int, clients chan counterClient, errors chan error, wg *sync.WaitGroup) {
Loop:
	for {
		select {
		// TODO: Needs timeouts...
		case c, ok := <-clients:
			if ok {
				val := service.CounterValue{int32(count)}
				if err := c.Send(&val); err != nil {
					errors <- err
				}
			} else {
				wg.Done()
				break Loop
			}
		}
	}
}
