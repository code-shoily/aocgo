package day01

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	var calories []int

	for _, cluster := range strings.Split(input, "\n\n") {
		var elfCalorie int
		for _, line := range strings.Split(cluster, "\n") {
			if calorie, error := strconv.Atoi(line); error == nil {
				elfCalorie += calorie
			} else {
				panic("Invalid data")
			}
		}

		calories = append(calories, elfCalorie)
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories[0], calories[0] + calories[1] + calories[2]
}
