package server

import "sync/atomic"

type counter int32

func (c *counter) increment(val int32) int32 {
	return atomic.AddInt32((*int32)(c), val)
}
