package main

import (
	"fmt"
	"sort"
)

// binarySearch is a function that search item in sorted array and returns his position, otherwise -1
func binarySearch(array []int, target int) int {
	min, max := 0, len(array)-1

	for min <= max {
		mid := min + ((max - min) / 2)

		switch {
		case array[mid] == target:
			return mid
		case array[mid] > target:
			max = mid - 1
		case array[mid] < target:
			min = mid + 1
		}
	}

	return -1
}

func main() {
	sortedArray := []int{-2, -1, 0, 1, 2, 3, 4, 5, 10, 20, 50, 100, 120}

	fmt.Println(binarySearch(sortedArray, 3))
	fmt.Println(sort.SearchInts(sortedArray, 3))
}
