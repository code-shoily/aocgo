// Package seq has helpers to manipulate sequences of data.
package seq

import "golang.org/x/exp/constraints"

// ChunkBy chunks `seq` into integer of size `by`.
func ChunkBy[T any](seq []T, by int) (chunks [][]T) {
	// FIXME: This only takes care of slices with lengths divisible by the chunk size.
	for i := 0; i < len(seq); i += by {
		chunk := make([]T, by)
		for j := 0; j < by; j++ {
			chunk[j] = seq[i+j]
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}

// Frequencies returns a map containing the frequencies of elements of a slice
func Frequencies[T comparable](sequence []T) map[T]int {
	frequencies := make(map[T]int)

	for _, value := range sequence {
		frequencies[value]++
	}

	return frequencies
}

// GetMinMax returns the smallest and largest value from the list.
func GetMinMax[T constraints.Ordered](list []T) (T, T) {
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

// Transpose transposes a 2D slice (from MxN to NxM).
func Transpose[T any](data [][]T) [][]T {
	transposed := make([][]T, len(data[0]))

	for idx := range transposed {
		transposed[idx] = make([]T, len(data))
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			transposed[j][i] = data[i][j]
		}
	}

	return transposed
}

// MakeSet turns a slice into set emulating map
func MakeSet[T comparable](seq []T) map[T]bool {
	set := make(map[T]bool)

	for _, elem := range seq {
		set[elem] = true
	}

	return set
}
