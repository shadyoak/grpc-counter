package client

import (
	"errors"
	"fmt"
	"io"

	"github.com/shadyoak/grpc-counter/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CounterClient struct {
	cleanup  func() error
	hostname string
	port     int
	Receive  ReceiveChannels
	stream   service.Counter_IncrementCounterClient
}

type ReceiveChannels struct {
	CounterUpdate chan int
	Error         chan error
	Done          chan bool
}

var ClientNotStartedError = errors.New("client not started")

func createCleanupFunc(stream service.Counter_IncrementCounterClient, conn *grpc.ClientConn) func() error {
	return func() error {
		streamErr := stream.CloseSend()
		connErr := conn.Close()
		if streamErr != nil {
			return streamErr
		}
		return connErr
	}
}

func (c *CounterClient) listenForStreamUpdates(stream service.Counter_IncrementCounterClient) <-chan bool {
	ch := make(chan bool)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				c.Receive.Done <- true
				return
			} else if err != nil {
				c.Receive.Error <- err
			} else {
				count := int(in.Count)
				c.Receive.CounterUpdate <- count
			}
		}
	}()
	return ch
}

func (c *CounterClient) openIncrementCounterStream(client service.CounterClient) (service.Counter_IncrementCounterClient, error) {
	stream, err := client.IncrementCounter(context.Background())
	if err != nil {
		return nil, err
	}
	c.listenForStreamUpdates(stream)
	return stream, nil
}

func (c *CounterClient) Address() string {
	return fmt.Sprintf("%v:%v", c.hostname, c.port)
}

func (c *CounterClient) IncrementCounter(val int) error {
	if c.stream == nil {
		return ClientNotStartedError
	}
	counter := &service.CounterValue{int32(1)}
	return c.stream.Send(counter)
}

func New(hostname string, port int) *CounterClient {
	receive := ReceiveChannels{
		CounterUpdate: make(chan int),
		Error:         make(chan error),
		Done:          make(chan bool),
	}
	return &CounterClient{
		hostname: hostname,
		port:     port,
		Receive:  receive,
	}
}

func (c *CounterClient) Start() error {

	conn, err := grpc.Dial(c.Address(), grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := service.NewCounterClient(conn)
	stream, err := c.openIncrementCounterStream(client)
	if err != nil {
		conn.Close()
		return err
	}

	c.cleanup = createCleanupFunc(stream, conn)
	c.stream = stream
	return nil
}

func (c *CounterClient) Stop() error {
	var err error
	if c.cleanup != nil {
		err = c.cleanup()
		c.stream = nil
		c.cleanup = nil
	}
	return err
}
