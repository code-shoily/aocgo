package algo

import (
	"math"
)

// Factors returns a list of factorials for given number
func Factors(number int) (factorials []int) {
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if remainder := number % i; remainder == 0 {
			division := number / i
			factorials = append(factorials, i)
			if division != i {
				factorials = append(factorials, division)
			}
		}
	}

	return factorials
}
