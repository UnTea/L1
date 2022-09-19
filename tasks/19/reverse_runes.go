package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	stringToReverse := "人民日报 главрыба 网是中国领先的国"

	fmt.Println(stringToReverse)
	fmt.Println(reverseString(stringToReverse))
}
