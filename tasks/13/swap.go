package main

import "fmt"

func main() {
	first, second := "first", "second"

	fmt.Println(first, second)

	first, second = second, first

	fmt.Println(first, second)
}
