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
	mutex   sync.RWMutex
}

type sendError struct {
	client *counterClient
	err    error
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
	errChan := make(chan sendError)
	errorList := make([]sendError, 0)

	wg := sync.WaitGroup{}
	wg.Add(4)
	go sendCount(count, clients, errChan, &wg)
	go sendCount(count, clients, errChan, &wg)
	go sendCount(count, clients, errChan, &wg)
	go sendCount(count, clients, errChan, &wg)
	go processErrors(errChan, &errorList)

	c.mutex.RLock()
	for _, c := range c.clients {
		clients <- c
	}
	close(clients)
	c.mutex.RUnlock()

	wg.Wait()
	close(errChan)

	for _, err := range errorList {
		if err.err == NotificationTimeoutError || *err.client == client {
			return err.err
		}
	}

	return nil
}

func processErrors(errChan chan sendError, acc *[]sendError) {
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

func sendCount(count int, clients chan counterClient, errChan chan sendError, wg *sync.WaitGroup) {
Loop:
	for {

		timeout := time.After(25 * time.Millisecond)

		select {
		case c, ok := <-clients:
			if ok {
				val := service.CounterValue{int32(count)}
				if err := c.Send(&val); err != nil {
					errChan <- sendError{client: &c, err: err}
				}
			} else {
				wg.Done()
				break Loop
			}
		case <-timeout:
			errChan <- sendError{err: NotificationTimeoutError}
		}
	}
}
