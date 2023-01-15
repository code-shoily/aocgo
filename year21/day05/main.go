// Package day05 - Solution for Advent of Code 2021/05
// Problem Link: http://adventofcode.com/2021/day/05
package day05

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	lines := parse(input)
	points := map[Point]int{}
	return solvePart1(lines, points), solvePart2(lines, points)
}

func solvePart1(lines []Line, points map[Point]int) int {
	for _, line := range lines {
		switch {
		case line.isVertical():
			for _, point := range line.getVerticalPoints() {
				points[point]++
			}
		case line.isHorizontal():
			for _, point := range line.getHorizontalPoints() {
				points[point]++
			}
		}
	}

	return countOverlaps(points)
}

func solvePart2(lines []Line, points map[Point]int) int {
	for _, line := range lines {
		if line.isDiagonal() {
			for _, point := range line.getDiagonalPoints() {
				points[point]++
			}
		}
	}

	return countOverlaps(points)
}

func parse(input string) (data []Line) {
	for _, line := range strings.Split(input, "\n") {
		data = append(data, parseLine(line))
	}

	return data
}

func parseLine(line string) Line {
	var x1, y1, x2, y2 int
	fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

	if x1 > x2 {
		return Line{Point{x2, y2}, Point{x1, y1}}
	}
	return Line{Point{x1, y1}, Point{x2, y2}}
}

func countOverlaps(points map[Point]int) (total int) {
	for _, v := range points {
		if v > 1 {
			total++
		}
	}

	return total
}

type Point = [2]int

type Line struct {
	from Point
	to   Point
}

func (line Line) isVertical() bool {
	return line.from[1] == line.to[1]
}

func (line Line) isHorizontal() bool {
	return line.from[0] == line.to[0]
}

func (line Line) isDiagonal() bool {
	return !line.isVertical() && !line.isHorizontal()
}

func (line Line) getHorizontalPoints() (points []Point) {
	yFrom, yTo := line.from[1], line.to[1]
	if yFrom > yTo {
		yTo, yFrom = yFrom, yTo
	}

	for y := yFrom; y <= yTo; y++ {
		points = append(points, Point{line.from[0], y})
	}

	return points
}

func (line Line) getVerticalPoints() (points []Point) {
	xFrom, xTo := line.from[0], line.to[0]
	if xFrom > xTo {
		xTo, xFrom = xFrom, xTo
	}

	for x := xFrom; x <= xTo; x++ {
		points = append(points, Point{x, line.from[1]})
	}

	return points
}

func (line Line) getDiagonalPoints() (points []Point) {
	xFrom, xTo, yFrom, yTo := line.from[0], line.to[0], line.from[1], line.to[1]

	switch {
	case xFrom == yFrom:
		if xFrom > xTo {
			xFrom, xTo = xTo, xFrom
		}

		for i := xFrom; i <= xTo; i++ {
			points = append(points, Point{i, i})
		}
	case xFrom <= xTo && yFrom >= yTo:
		for xFrom <= xTo {
			points = append(points, Point{xFrom, yFrom})
			xFrom++
			yFrom--
		}
	case xFrom <= xTo && yFrom <= yTo:
		for xFrom <= xTo {
			points = append(points, Point{xFrom, yFrom})
			xFrom++
			yFrom++
		}
	case xFrom >= xTo && yFrom <= yTo:
		for xFrom >= xTo {
			points = append(points, Point{xFrom, yFrom})
			xFrom--
			yFrom++
		}
	case xFrom >= xTo && yFrom >= yTo:
		for xFrom >= xTo {
			points = append(points, Point{xFrom, yFrom})
			xFrom--
			yFrom--
		}
	}

	return points
}
