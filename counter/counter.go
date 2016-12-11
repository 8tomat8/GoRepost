package counter

import (
	"sync"
	"sync/atomic"
)

var instance *Counter
var once sync.Once

// GetCounter returns Singleton Counter object
func GetCounter() *Counter {
	once.Do(func() {
		instance = &Counter{Count: 0}
	})
	return instance
}

// Counter of tasks in processing
type Counter struct {
	Count uint64
}

func (c Counter) JobStarted() {
	atomic.AddUint64(&c.Count, 1)
}

func (c Counter) JobFinished() {
	atomic.AddUint64(&c.Count, ^uint64(0))
}

func (c Counter) GetSize() uint64 {
	return atomic.LoadUint64(&c.Count)
}
