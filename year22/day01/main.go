package day01

import (
	"container/heap"
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/io"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	calorieHeap := parse(input)
	heap.Init(calorieHeap)

	maxCalorie := calorieHeap.Max().(int)
	top3CalorieSum := heap.Pop(calorieHeap).(int) +
		heap.Pop(calorieHeap).(int) + heap.Pop(calorieHeap).(int)

	return maxCalorie, top3CalorieSum
}

func parse(input string) *algo.MaxHeap[int] {
	clusters := strings.Split(input, "\n\n")
	calorieHeap := make(algo.MaxHeap[int], 0, len(clusters))

	for _, cluster := range clusters {
		var elfCalorie int
		for _, calorie := range io.SplitIntLines(cluster) {
			elfCalorie += calorie
		}
		calorieHeap.Push(elfCalorie)
	}

	return &calorieHeap
}
