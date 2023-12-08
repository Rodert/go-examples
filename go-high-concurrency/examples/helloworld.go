package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	v   int
	mux sync.Mutex
}

func (c *Counter) Inc() {
	c.mux.Lock()
	c.v++
	c.mux.Unlock()
}

func (c *Counter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v
}

func main() {
	var wg sync.WaitGroup
	var counter Counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
