package challenges

import (
	"fmt"
	"golang-essential-training-course/03-object-oriented/structs"
	"log"
)

func SquareChallenge() {
	s, err := structs.NewSquare(0, 0, 10)

	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	s.Move(15, 10)
	fmt.Printf("New center of square = %+v\n", s.GetCenter())

	fmt.Printf("area of square = %d", s.Area())
}
