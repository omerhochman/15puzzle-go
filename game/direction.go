package game

import "strings"

type Direction string

func (dir Direction) String() string {
	return string(dir)
}

func DirectionFromString(str string) Direction {
	direction := strings.ToLower(str)

	switch direction {
	case string(UP):
		return UP
	case string(DOWN):
		return DOWN
	case string(LEFT):
		return LEFT
	case string(RIGHT):
		return RIGHT
		return Direction(direction)
	default:
		return ""
	}
}
