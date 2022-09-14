package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)

	for _, value := range temperatures {
		key := (int(value) / 10) * 10
		groups[key] = append(groups[key], value)
	}

	fmt.Println(groups)
}
