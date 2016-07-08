package server

import (
	"log"
	"sync"
	"time"

	"github.com/shadyoak/grpc-counter/service"
)

type counterClient interface {
	Send(m *service.CounterValue) error
}

type clientList struct {
	clients []counterClient
	mutex   sync.Mutex
}

func (c *clientList) addClient(client counterClient) {
	c.mutex.Lock()
	c.clients = append(c.clients, client)
	c.mutex.Unlock()
	log.Println("client connected:", client)
}

func newClientList() clientList {
	c := []counterClient{}
	m := sync.Mutex{}
	return clientList{c, m}
}

func (c *clientList) notifyAllClients(client counterClient, count int) error {

	clients := make(chan counterClient)
	errChan := make(chan error)
	errorList := make([]error, 0)

	wg := sync.WaitGroup{}
	wg.Add(4)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go processErrors(errChan, &errorList)

	c.mutex.Lock()
	for _, c := range c.clients {
		clients <- c
	}
	close(clients)

	wg.Wait()
	close(errChan)

	c.mutex.Unlock()

	if len(errorList) > 0 {
		return errorList[0]
	}

	return nil
}

func processErrors(errChan chan error, acc *[]error) {
Loop:
	for {
		select {
		case err, ok := <-errChan:
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

func sendCount(count int, curr counterClient, clients chan counterClient, errChan chan error, wg *sync.WaitGroup) {
Loop:
	for {

		timeout := time.After(2000 * time.Millisecond)

		select {
		case c, ok := <-clients:
			if ok {
				val := service.CounterValue{int32(count)}
				if err := c.Send(&val); err != nil && c == curr {
					errChan <- err
				}
			} else {
				wg.Done()
				break Loop
			}
		case <-timeout:
			errChan <- NotificationTimeoutError
		}
	}
}
