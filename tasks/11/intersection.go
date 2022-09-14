package main

import "fmt"

func Intersection(firstSet, secondSet []int) []int {
	setA := make(map[int]bool, len(firstSet))
	setB := make(map[int]bool, len(secondSet))

	var result []int

	for _, value := range firstSet {
		setA[value] = true
	}

	for _, value := range secondSet {
		setB[value] = true
	}

	for key := range setB {
		if setA[key] {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	setA := []int{-5, 2, -3, 4, -5, 6, -7, 8}
	setB := []int{4, -5, -6, 8, -1, -10, 9, 0, 6, -6, 7, -8}

	fmt.Println(Intersection(setA, setB))

	setA = []int{-7, 3, 5, -5, 7, -4, -9, -3, 8, -10, -2, -8, 6, 2, 10, 9, -6, -1, 4, 1}
	setB = []int{1, -1, -3, 4, 7, 10, -5, -9, 0, -2, -4, 9, 6, 5, 3, -10, -6, 2, 8, -7, -8}

	fmt.Println(Intersection(setA, setB))
}
