// Package seq has helpers to manipulate sequences of data.
package seq

import "golang.org/x/exp/constraints"

// Chunk chunks `seq` into integer of `size`. `interval` decides overlap, when `discard` is true, the remainder
// is discarded if not of equal size.
func Chunk[T any](seq []T, size int, interval int, discard bool) (chunks [][]T) {
	i, j := 0, size

	for j < len(seq) {
		chunk := make([]T, size)
		copy(chunk, seq[i:j])
		chunks = append(chunks, chunk)
		i += interval
		j += interval
	}

	if lastChunk := seq[i:]; len(chunks[0]) == len(lastChunk) || !discard {
		chunks = append(chunks, lastChunk)
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

// Reverse reverses a slice
func Reverse[T any](seq []T) []T {
	reversed := make([]T, len(seq))
	for idx, elem := range seq {
		reversed[len(seq)-1-idx] = elem
	}

	return reversed
}
