package game

import "strings"

type Direction string

func (dir Direction) String() string {
	return string(dir)
}

func DirectionFromString(str string) Direction {
	directionAsText := strings.ToLower(str)

	directions := map[string]Direction{
		"up":    UP,
		"down":  DOWN,
		"right": RIGHT,
		"left":  LEFT,
	}

	direction := directions[directionAsText]

	return direction
}
