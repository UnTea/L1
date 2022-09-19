package main

import (
	"fmt"
	"math/big"
)

// Sum is a fucntion that sums two numbers and returns the result
func Sum(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Add(lhs, rhs)
}

// Subtraction is a fucntion that subtractions two numbers and returns the result
func Subtraction(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Sub(lhs, rhs)
}

// Multiply is a fucntion that multiplies two numbers and returns the result
func Multiply(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Mul(lhs, rhs)
}

// Divide is a fucntion that divides two numbers and returns the result
func Divide(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Quo(lhs, rhs)
}

func main() {
	a := big.NewFloat(float64(1<<234) + float64(1<<219))
	b := big.NewFloat(float64(1<<233) + float64(1<<245))

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Sum is equal: ", Sum(a, b))
	fmt.Println("Subtraction is equal: ", Subtraction(a, b))
	fmt.Println("Multiply is equal: ", Multiply(a, b))
	fmt.Println("Division is equal: ", Divide(a, b))
}
