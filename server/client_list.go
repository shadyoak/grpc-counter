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
	clients map[counterClient]sync.Mutex
	mutex   sync.RWMutex
}

type clientMutexPair struct {
	client counterClient
	mutex  sync.Mutex
}

func (c *clientList) addClient(client counterClient) {
	c.mutex.Lock()
	c.clients[client] = sync.Mutex{}
	c.mutex.Unlock()
	log.Println("client connected:", client)
}

func newClientList() clientList {
	c := make(map[counterClient]sync.Mutex)
	m := sync.RWMutex{}
	return clientList{c, m}
}

func (c *clientList) notifyAllClients(client counterClient, count int) error {

	clients := make(chan clientMutexPair)
	errChan := make(chan error)
	errorList := make([]error, 0)

	wg := sync.WaitGroup{}
	wg.Add(4)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go sendCount(count, client, clients, errChan, &wg)
	go processErrors(errChan, &errorList)

	c.mutex.RLock()
	for k, v := range c.clients {
		clients <- clientMutexPair{client: k, mutex: v}
	}
	close(clients)
	c.mutex.RUnlock()

	wg.Wait()
	close(errChan)

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
	delete(c.clients, client)
	c.mutex.Unlock()
	log.Println("client disconnected:", client)
}

func sendCount(count int, curr counterClient, clients chan clientMutexPair, errChan chan error, wg *sync.WaitGroup) {
Loop:
	for {

		timeout := time.After(2000 * time.Millisecond)

		select {
		case pair, ok := <-clients:
			if ok {
				val := service.CounterValue{int32(count)}
				c := pair.client

				pair.mutex.Lock()
				if err := c.Send(&val); err != nil && c == curr {
					errChan <- err
				}
				pair.mutex.Unlock()

			} else {
				wg.Done()
				break Loop
			}
		case <-timeout:
			errChan <- NotificationTimeoutError
		}
	}
}
