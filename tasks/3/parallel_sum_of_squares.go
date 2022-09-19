package main

import (
	"fmt"
	"sync"
)

// powerInteger is a function that powered an integer number to a power
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

// SumOfSquaresWaitGroup a function that parallel calculates the sum of squares of the array elements using sync.WaitGroup
func SumOfSquaresWaitGroup(array []int) int {
	if len(array) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	sum := 0

	wg.Add(len(array))

	for i := 0; i < len(array); i++ {
		go func(index int) {
			defer wg.Done()

			sum += powerInteger(array[index], 2)
		}(i)
	}

	wg.Wait()

	return sum
}

// SumOfSquaresChannel a function that parallel calculates the sum of squares of the array elements using channels
func SumOfSquaresChannel(array []int) int {
	if len(array) == 0 {
		return 0
	}

	channel := make(chan int)
	sum := 0

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

		sum += value
	}

	return sum
}

func main() {
	numbers := [...]int{2, 4, 6, 8, 10}
	sumOfSquaresWaitGroup := SumOfSquaresWaitGroup(numbers[:])

	fmt.Printf("SumOfSquaresWaitGroup result: %v\n", sumOfSquaresWaitGroup)

	sumOfSquaresChannel := SumOfSquaresChannel(numbers[:])

	fmt.Printf("SumOfSquaresChannel result: %v\n", sumOfSquaresChannel)
}
