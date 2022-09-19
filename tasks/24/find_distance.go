package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("\n x: %f y: %f\n", p.x, p.y)
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Distance(lhs, rhs *Point) float64 {
	return math.Sqrt(math.Pow(rhs.GetX()-lhs.GetX(), 2) + math.Pow(rhs.GetY()-lhs.GetY(), 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(100, 100)

	fmt.Printf("p1: %vp2: %v", p1, p2)
	fmt.Println("Distance between two point is: ", Distance(p1, p2))

	p1 = NewPoint(22, 22)
	p2 = NewPoint(88, 88)

	fmt.Printf("\np1: %vp2: %v", p1, p2)
	fmt.Println("Distance between two point is: ", Distance(p1, p2))
}
