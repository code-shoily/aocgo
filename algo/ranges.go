package algo

// GetMinMax returns the smallest and largest value from the list.
func GetMinMax(list []int) (int, int) {
	min, max := list[0], list[0]

	for _, elem := range list {
		if elem > max {
			max = elem
		}
		if elem < min {
			min = elem
		}
	}

	return min, max
}
