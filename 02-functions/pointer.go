package _2_functions

func DoubleUsage(number int) int {
	double(&number)
	return number
}

func double(number *int) {
	*number *= 2
}
