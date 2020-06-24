package challenges

import (
	"fmt"
	"strings"
)

func CountWords() {
	text := `
		Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind
	`

	words := strings.Fields(text)
	wordsCount := map[string]int{}

	for _, word := range words {
		wordsCount[strings.ToLower(word)]++
	}

	fmt.Println(wordsCount)
}
