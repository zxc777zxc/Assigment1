package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

type Triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (s Square) Perimeter() float64 {
	return s.length * 4
}

func (t Triangle) Area() float64 {
	p := (t.sideA + t.sideB + t.sideC) / 2.0
	return math.Sqrt(p * (p - t.sideA) * (p - t.sideB) * (p - t.sideC))
}

func (t Triangle) Perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("Shape details:\nArea: %.2f\nPerimeter: %.2f\n\n", s.Area(), s.Perimeter())
}
