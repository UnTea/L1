package main

import (
	"fmt"
	"sync"
)

// Square is a function that squares values from input channel and sen then to output channel
func Square(input <-chan int, output chan<- int) {
	for value := range input {
		square := value * value
		output <- square
	}

	close(output)
}

// Read is a function that reads data from input channel and prints it
func Read(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range input {
		fmt.Println(value)
	}
}

func main() {
	input, output := make(chan int, 10), make(chan int, 10)
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go Square(input, output)

	var wg sync.WaitGroup

	wg.Add(1)
	go Read(output, &wg)

	for _, value := range array {
		input <- value
	}

	close(input)

	wg.Wait()
}
