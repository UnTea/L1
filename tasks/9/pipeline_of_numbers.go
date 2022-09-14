package main

import (
	"fmt"
	"sync"
)

func Square(in <-chan int, out chan<- int) {
	for val := range in {
		square := val * val
		out <- square
	}

	close(out)
}

func Read(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	in, out := make(chan int, 10), make(chan int, 10)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go Square(in, out)

	var wg sync.WaitGroup

	wg.Add(1)
	go Read(out, &wg)

	for _, val := range arr {
		in <- val
	}
	close(in)

	wg.Wait()
}
