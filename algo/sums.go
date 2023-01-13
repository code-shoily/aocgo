package algo

// TwoSum returns the two values of an int slice that sums to target.
func TwoSum(sortedData []int, target int) (int, int, bool) {
	left := 0
	right := len(sortedData) - 1

	for left < right {
		currentTotal := sortedData[left] + sortedData[right]
		switch {
		case currentTotal > target:
			right--
		case currentTotal < target:
			left++
		case currentTotal == target:
			return sortedData[left], sortedData[right], true
		}
	}

	return -1, -1, false
}

// ThreeSum returns three values of a slice that sums to target.
func ThreeSum(sortedData []int, target int) (int, int, int, bool) {
	for current := 0; current < len(sortedData); current++ {
		currentValue := sortedData[current]
		remaining := target - currentValue

		if a, b, ok := TwoSum(sortedData[current+1:], remaining); ok {
			return currentValue, a, b, true
		}
	}

	return -1, -1, -1, false
}
