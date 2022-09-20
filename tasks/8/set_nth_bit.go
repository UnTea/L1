package main

import (
	"fmt"
	"strconv"
)

func ChangeNthBit(number *int64, position int64, value int64) {
	position -= 1                       // reduction to the zero-based numbering form
	check := *number&(1<<position) != 0 // checking for the presence of a bit

	if check && (value == 1) || !check && (value == 0) { // checking cases when the bit is already in the desired position
		return
	} else if !check && (value == 1) { // the value of bit 0 must be set to 1
		*number |= 1 << position
	} else {
		*number &= ^(1 << position) // the value of bit 1 must be set to 0
	}
}

func ChangeNthBitByMasks(number *int64, position int64, value int64) {
	position -= 1                // reduction to the zero-based numbering form
	*number &= ^(1 << position)  // zeroed a bit in position
	*number |= value << position // bitwise or for valued a bit at position
}

func main() {
	number := int64(43690)

	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010101010101010 43690

	ChangeNthBit(&number, 11, 1)
	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010111010101010 44714

	ChangeNthBitByMasks(&number, 2, 0)
	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010111010101000 44712

	ChangeNthBitByMasks(&number, 4, 0)
	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010111010100000 44704

	ChangeNthBitByMasks(&number, 2, 1)
	fmt.Println(strconv.FormatInt(number, 2), strconv.FormatInt(number, 10)) // 1010111010100010 44706
}
