package algo

func IntToDigits(n int) []int {
	var reverseDigits []int
	for n > 0 {
		reverseDigits = append(reverseDigits, n%10)
		n /= 10
	}

	digitSize := len(reverseDigits)
	digits := make([]int, digitSize)

	for idx, digit := range reverseDigits {
		digits[digitSize-1-idx] = digit
	}

	return digits
}
