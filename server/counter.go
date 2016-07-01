package server

import "sync/atomic"

type Counter int32

func (c *Counter) Increment(val int32) int32 {
	return atomic.AddInt32((*int32)(c), val)
}

func (c *Counter) Value() int32 {
	return atomic.LoadInt32((*int32)(c))
}
