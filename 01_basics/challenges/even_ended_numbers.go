package challenges

import "fmt"

func CountEvenEndedNumbers() int {
	count := 0

	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {

			product := intToString(a * b)

			if firstEqualsLast(product) {
				count++
			}
		}
	}

	return count
}

func intToString(number int) string {
	return fmt.Sprintf("%d", number)
}

func firstEqualsLast(number string) bool {
	return number[0] == number[len(number)-1]
}
