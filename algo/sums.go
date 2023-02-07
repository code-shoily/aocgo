package algo

import (
	"github.com/code-shoily/aocgo/seq"
	"golang.org/x/exp/constraints"
	"sort"
)

// TwoSumSorted returns the two values of a sorted number slice that sums to target.
func TwoSumSorted[N constraints.Integer](sortedData []N, target N) (N, N, bool) {
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

	return 0, 0, false
}

// ThreeSumSorted returns three values of a sorted number slice that sums to target.
func ThreeSumSorted[N constraints.Integer](sortedData []N, target N) (N, N, N, bool) {
	for current := 0; current < len(sortedData); current++ {
		currentValue := sortedData[current]
		remaining := target - currentValue

		if a, b, ok := TwoSumSorted[N](sortedData[current+1:], remaining); ok {
			return currentValue, a, b, true
		}
	}

	return 0, 0, 0, false
}

// TwoSum returns the two values of a number slice that sums to target.
func TwoSum[N constraints.Integer](data []N, target N) (N, N, bool) {
	sortedData := make([]N, len(data))
	copy(sortedData, data)
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i] < sortedData[j]
	})

	return TwoSumSorted(sortedData, target)
}

// ThreeSum returns three values of a number slice that sums to target.
func ThreeSum[N constraints.Integer](data []N, target N) (N, N, N, bool) {
	sortedData := make([]N, len(data))
	copy(sortedData, data)
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i] < sortedData[j]
	})

	return ThreeSumSorted(sortedData, target)
}

// SubArraySum returns the subarray that sums to given target
func SubArraySum[N constraints.Integer](data []N, target N) (from, to int, found bool) {
	memo := map[N]int{}
	var runningSum N
	for currentIdx, v := range data {
		runningSum += v

		if runningSum == target {
			return 0, currentIdx, true
		}

		if start, exists := memo[runningSum-target]; exists {
			return start + 1, currentIdx, true
		}

		memo[runningSum] = currentIdx
	}

	return
}

// SubsetSum returns a list of elements that belong to a list that sum to target
func SubsetSum[T constraints.Integer](numbers []T, target T) (result [][]T) {
	subsetSumUtil(numbers, target, []T{}, &result)
	return result
}

func subsetSumUtil[T constraints.Integer](numbers []T, target T, partial []T, result *[][]T) {
	s := seq.Sum(partial)

	if s == target {
		*result = append(*result, partial)
	}
	if s >= target {
		return
	}

	for i := range numbers {
		n := numbers[i]
		remaining := numbers[i+1:]
		subsetSumUtil(remaining, target, append(partial, n), result)
	}
}
