package _2_functions

import (
	"fmt"
	"math"
)

func SqrtUsage(number float64) {
	if sqrt, err := sqrt(number); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(sqrt)
	}
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", number)
	}

	return math.Sqrt(number), nil
}
