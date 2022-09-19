package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)

	left, right := 0, len(words)-1

	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

func reverseWords1(str string) string {
	words := strings.Fields(str)

	var result strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		result.WriteString(" ")
	}

	return strings.TrimSpace(result.String())
}

func main() {
	stringToReverse := "snow dog sun"

	fmt.Println(stringToReverse)
	fmt.Println(reverseWords(stringToReverse))
	fmt.Println(reverseWords1(stringToReverse))
}
