package main

import "fmt"

func main() {
	first, second := 1, 2

	fmt.Println(first, second)

	first, second = second, first

	fmt.Println(first, second)
}
