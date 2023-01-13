package day01

type Facing int

const (
	North Facing = iota
	South
	East
	West
)

var FaceFactors = map[Facing]Point{
	North: Point{0, 1},
	South: Point{0, -1},
	East:  Point{1, 0},
	West:  Point{-1, 0},
}

func (facing Facing) onLeft() Facing {
	switch facing {
	case North:
		return West
	case South:
		return East
	case East:
		return North
	case West:
		return South
	default:
		panic("Impossible to reach here")
	}
}

func (facing Facing) onRight() Facing {
	switch facing {
	case North:
		return East
	case South:
		return West
	case East:
		return South
	case West:
		return North
	default:
		panic("Impossible to reach here")
	}
}

func (facing Facing) navigate(direction string) Facing {
	if direction == "L" {
		return facing.onLeft()
	} else {
		return facing.onRight()
	}
}
