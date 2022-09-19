package main

import (
	"fmt"
	"sync"
)

// Counter is a structure that counters with mutex for concurrent counting
type Counter struct {
	sync.Mutex
	counter int
}

// NewCounter is a Counter constructor
func NewCounter() *Counter {
	return &Counter{}
}

// Increment is a function that increments the counter field thread-safely
func (c *Counter) Increment() {
	c.Lock()
	c.counter++
	c.Unlock()
}

// String is a function that prints the counter field
func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func main() {
	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
