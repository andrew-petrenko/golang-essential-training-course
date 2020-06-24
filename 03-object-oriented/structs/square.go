package structs

import (
	"fmt"
)

type square struct {
	center Point
	length int
}

func NewSquare(x, y, length int) (*square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)", length)
	}

	s := square{
		Point{x, y},
		length,
	}

	return &s, nil
}

func (s *square) GetCenter() *Point {
	return &s.center
}

func (s *square) GetLength() *int {
	return &s.length
}

func (s *square) Move(dx, dy int) {
	s.GetCenter().Move(dx, dy)
}

func (s *square) Area() int {
	return s.length * s.length
}
