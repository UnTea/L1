package main

import "fmt"

// Set is a function that returns set of unique elements of the input set
func Set(slice []string) []string {
	var result []string
	set := make(map[string]bool, len(slice))

	for _, value := range slice {
		set[value] = true
	}

	for key := range set {
		result = append(result, key)
	}

	return result
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set(slice)

	fmt.Printf("Slice representation: %v\nSet   representation: %v\n", slice, set)
}
