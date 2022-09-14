package main

import (
	"fmt"
	"strconv"
)

func ChangeNthBit(number *int64, position uint, value uint) {
	check := *number&(1<<int64(position)) != 0 // checking for the presence of a bit

	if check && (value == 1) || !check && (value == 0) { // checking cases when the bit is already in the desired position
		return
	} else if !check && (value == 1) { // the value of bit 0 must be set to 1
		*number |= 1 << position
	} else {
		*number &= ^(1 << position) // the value of bit 1 must be set to 0
	}
}

func main() {
	number := int64(43690)

	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010101010101010 43690
	ChangeNthBit(&number, 10, 1)
	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010111010101010 44714
}
