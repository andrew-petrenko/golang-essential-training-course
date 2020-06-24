package challenges

func MaxValue() int {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]

	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	return max
}
