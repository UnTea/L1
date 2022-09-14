package main

import (
	"fmt"
	"sync"
)

type Map struct {
	sync.Mutex
	sync.WaitGroup
	m map[int]int
}

func NewMap() *Map {
	return &Map{
		m: make(map[int]int),
	}
}

func (m *Map) SetValue(id int) {
	m.Lock()
	for i := 0; i <= 10; i++ {
		m.m[i] = id
	}
	m.Unlock()

	m.Done()
}

func main() {
	m := NewMap()

	m.Add(10)

	go m.SetValue(1)
	go m.SetValue(2)
	go m.SetValue(3)
	go m.SetValue(4)
	go m.SetValue(5)
	go m.SetValue(6)
	go m.SetValue(7)
	go m.SetValue(8)
	go m.SetValue(9)
	go m.SetValue(10)

	m.Wait()

	fmt.Println(m.m)
}
