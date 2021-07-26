package main

import (
	"math"
	"strings"
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
	name() string
}

type square struct {
	sideLength float64
}

func (s square) name() string {
	return "square"
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (s square) perimeter() float64 {
	return s.sideLength * 4
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) name() string {
	return "rectangle"
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func main() {
	s := square{10.0}
	r := rectangle{width: 5.0, height: 4.0}
	c := circle{1.0}

	shapes := []shape{s, r, c}

	for _, shape := range shapes {
		fmt.Println(strings.Title(shape.name()))
		fmt.Println("Area:", shape.area())
		fmt.Println("Perimiter", shape.perimeter())
		fmt.Println("")
	}
}
