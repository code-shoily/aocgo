package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	expenses := parse(input)
	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i] < expenses[j]
	})
	return solvePart1(expenses), solvePart2(expenses)
}

func solvePart1(expenses []int) int {
	if a, b, ok := TwoSum(expenses, 2020); ok {
		return a * b
	}

	panic("Nothing sums to 2020")
}

func solvePart2(expenses []int) int {
	for current := 0; current < len(expenses); current++ {
		currentExpense := expenses[current]
		remaining := 2020 - currentExpense

		if a, b, ok := TwoSum(expenses[current+1:], remaining); ok {
			return a * b * currentExpense
		}
	}

	panic("Nothing sums to 2020")
}

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

func parse(input string) (expenses []int) {
	for _, expense := range strings.Split(input, "\n") {
		if expenseValue, err := strconv.Atoi(expense); err == nil {
			expenses = append(expenses, expenseValue)
		}
	}

	return expenses
}
