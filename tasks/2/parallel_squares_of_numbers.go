package main

import (
	"fmt"
	"sync"
)

func powerInteger(number, power int) int {
	if power == 0 {
		return 1
	}

	result := number

	for i := 2; i <= power; i++ {
		result *= number
	}

	return result
}

func SquaresWaitGroup(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	var wg sync.WaitGroup

	squares := make([]int, len(array))

	wg.Add(len(array))

	for i := 0; i < len(array); i++ {
		go func(index int) {
			defer wg.Done()

			square := powerInteger(array[index], 2)

			squares[index] = square
		}(i)
	}

	wg.Wait()

	return squares
}

func SquaresChannel(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	channel := make(chan int)
	squares := make([]int, 0)

	go func() {
		for _, value := range array {
			channel <- powerInteger(value, 2)
		}

		close(channel)
	}()

	for {
		value, ok := <-channel
		if !ok {
			break
		}

		squares = append(squares, value)
	}

	return squares
}

func main() {
	numbers := [...]int{2, 4, 6, 8, 10}

	squaresWaitGroup := SquaresWaitGroup(numbers[:])
	fmt.Println("SquaresWaitGroup result:")

	for i := range numbers {
		fmt.Printf("\tsquare of %v is %v\n", numbers[i], squaresWaitGroup[i])
	}

	squaresChannel := SquaresChannel(numbers[:])
	fmt.Println("\nSquaresChannel result:")

	for i := range numbers {
		fmt.Printf("\tsquare of %v is %v\n", numbers[i], squaresChannel[i])
	}
}
