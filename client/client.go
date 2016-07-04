package client

import (
	"fmt"
	"io"
	"time"

	"github.com/shadyoak/grpc-counter/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type CounterClient struct {
	cleanup  func() error
	hostname string
	port     int
	Receive  ReceiveChannels
}

type ReceiveChannels struct {
	CounterUpdate chan int
	Error         chan error
	Done          chan bool
}

func (c *CounterClient) countToTen(stream service.Counter_IncrementCounterClient) <-chan bool {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			counter := &service.CounterValue{int32(1)}
			if err := stream.Send(counter); err != nil {
				// TODO: make non-fatal
				grpclog.Fatalf("failed to send a count: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}

func createCleanupFunc(stream service.Counter_IncrementCounterClient, conn *grpc.ClientConn) func() error {
	return func() error {
		err1 := stream.CloseSend()
		err2 := conn.Close()
		if err1 != nil {
			return err1
		} else if err2 != nil {
			return err2
		}
		return nil
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
	c.countToTen(stream)
	return stream, nil
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

	address := fmt.Sprintf("%v:%v", c.hostname, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
	return nil
}

func (c *CounterClient) Stop() error {
	var err error
	if c.cleanup != nil {
		err = c.cleanup()
		c.cleanup = nil
	}
	return err
}
