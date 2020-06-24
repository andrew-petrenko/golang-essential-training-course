package geometry

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
}

func (s *square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func sumAreas(shapes []shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.area()
	}

	return total
}

func SumOfShapes() {
	s := &square{20}
	fmt.Printf("Square area: %f\n", s.area())

	c := &circle{10}
	fmt.Printf("Circle area: %f\n", c.area())

	shapes := []shape{s, c}
	fmt.Printf("Sum of shape areas: %f\n", sumAreas(shapes))
}
