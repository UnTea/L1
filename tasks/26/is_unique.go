package main

import (
	"fmt"
	"strings"
)

func IsUnique(s string) bool {
	alpha := [26]int{}
	s = strings.ToLower(s)

	for _, v := range s {
		alpha[v-'a']++
	}

	for _, val := range alpha {
		if val > 1 {
			return false
		}
	}

	return true
}

func main() {
	array := [...]string{
		"abcd",
		"Abcd",
		"ABCDEFG",
		"abCdefAaf",
		"aabcd",
		"abcdeffg",
	}

	for i := range array {
		fmt.Printf("%s is %v\n", array[i], IsUnique(array[i]))
	}
}
