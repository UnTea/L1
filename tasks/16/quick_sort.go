package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quickSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	left, right := 0, len(array)-1

	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int() % len(array)

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	quickSort(array[:left])
	quickSort(array[left+1:])

	return array
}

func main() {
	array := []int{3, 7, 2, 9, 0, 1, 6, 8, 4}

	fmt.Println(array)
	fmt.Println(quickSort(array))
	sort.Ints(array)
	fmt.Println(array)
}
