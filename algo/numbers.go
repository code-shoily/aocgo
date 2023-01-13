package algo

import "math"

// IntToDigits explodes a number into a slice of digits.
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

// ToDecimal converts a list of bits into decimal number
func ToDecimal(bits []int) (decimal int) {
	size := len(bits)
	for i := 0; i < len(bits); i++ {
		decimal += bits[i] * int(math.Pow(2, float64(size-i-1)))
	}
	return decimal
}
