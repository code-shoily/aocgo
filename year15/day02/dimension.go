package day02

import (
	"strconv"
	"strings"
)

type Dimension struct {
	height int
	width  int
	length int
}

func (dim Dimension) surfaceArea() int {
	return 2 * (dim.length*dim.height + dim.height*dim.width + dim.width*dim.length)
}

func (dim Dimension) smallestSides() (int, int) {
	if dim.length < dim.height {
		if dim.width < dim.height {
			return dim.length, dim.width
		}
		return dim.length, dim.height
	} else {
		if dim.width < dim.length {
			return dim.height, dim.width
		}
		return dim.height, dim.length
	}
}

func (dim Dimension) smallestSurfaceArea() int {
	x, y := dim.smallestSides()

	return x * y
}

func (dim Dimension) smallestPerimeter() int {
	x, y := dim.smallestSides()

	return 2 * (x + y)
}

func (dim Dimension) volume() int {
	return dim.height * dim.width * dim.length
}

func newDimension(line string) Dimension {
	var dims []int
	for _, token := range strings.Split(line, "x") {
		if dim, error := strconv.Atoi(token); error == nil {
			dims = append(dims, dim)
		}
	}

	return Dimension{dims[0], dims[1], dims[2]}
}
