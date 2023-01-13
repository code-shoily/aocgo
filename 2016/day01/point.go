package main

import "math"

type Point struct {
	x int
	y int
}

func (p Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p Point) multiply(factor int) Point {
	return Point{p.x * factor, p.y * factor}
}

func (p Point) manhattanDistance(other Point) int {
	return int(math.Abs(float64(p.x-other.x))) + int(math.Abs(float64(p.y-other.y)))
}
