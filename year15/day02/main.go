package day02

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	dimensions := parse(input)

	return solver(dimensions, calculateWrappingPaper), solver(
		dimensions,
		calculateRibbon,
	)
}

func solver(dimensions []Dimension, solveFn func(dimension Dimension) int) (total int) {
	for _, dimension := range dimensions {
		total += solveFn(dimension)
	}
	return total
}

func parse(input string) (dims []Dimension) {
	for _, line := range io.SplitLines(input) {
		dims = append(dims, newDimension(line))
	}

	return dims
}

func calculateWrappingPaper(dimension Dimension) int {
	return dimension.smallestSurfaceArea() + dimension.surfaceArea()
}

func calculateRibbon(dimension Dimension) int {
	return dimension.smallestPerimeter() + dimension.volume()
}
