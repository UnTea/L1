package main

import (
	"fmt"
	"math/big"
)

func Add(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Add(lhs, rhs)
}

func Subtraction(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Sub(lhs, rhs)
}

func Multiply(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Mul(lhs, rhs)
}

func Divide(lhs, rhs *big.Float) *big.Float {
	return new(big.Float).Quo(lhs, rhs)
}

func main() {
	a := big.NewFloat(float64(1<<234) + float64(1<<219))
	b := big.NewFloat(float64(1<<233) + float64(1<<245))

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Add is equal: ", Add(a, b))
	fmt.Println("Subtraction is equal: ", Subtraction(a, b))
	fmt.Println("Multiply is equal: ", Multiply(a, b))
	fmt.Println("Division is equal: ", Divide(a, b))
}
