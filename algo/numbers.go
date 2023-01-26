package algo

import (
	"golang.org/x/exp/constraints"
	"math"
)

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

// IsEven returns if a number is even
func IsEven[T constraints.Integer](num T) bool {
	return num&1 == 0
}

// IsOdd returns if a number is odd
func IsOdd[T constraints.Integer](num T) bool {
	return !IsEven(num)
}

// CountSetBits returns the number of 1-s in a number's binary representation
func CountSetBits[T constraints.Integer](num T) (total T) {
	for num > 0 {
		total += num & 1
		num >>= 1
	}
	return total
}
