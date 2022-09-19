package main

import (
	"fmt"
	"math"
)

// Point is a struct that contains x and y coordinates
type Point struct {
	x, y float64
}

// NewPoint is a constructor of Point
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// String is a function that prints the Point fields
func (p *Point) String() string {
	return fmt.Sprintf("\n x: %f y: %f\n", p.x, p.y)
}

// GetX is a function that returns x field from Point object
func (p *Point) GetX() float64 {
	return p.x
}

// GetY is a function that returns x field from Point object
func (p *Point) GetY() float64 {
	return p.y
}

// Distance is a function that returns distance between two Point objects
func Distance(lhs, rhs *Point) float64 {
	return math.Sqrt(math.Pow(rhs.GetX()-lhs.GetX(), 2) + math.Pow(rhs.GetY()-lhs.GetY(), 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(100, 100)

	fmt.Printf("p1: %vp2: %v", p1, p2)
	fmt.Println("Distance between two points is: ", Distance(p1, p2))

	p1 = NewPoint(22, 22)
	p2 = NewPoint(88, 88)

	fmt.Printf("\np1: %vp2: %v", p1, p2)
	fmt.Println("Distance between two points is: ", Distance(p1, p2))
}
