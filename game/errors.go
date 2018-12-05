package game

import "fmt"

type CannotMoveCellError struct {
	number    int
	direction Direction
}

func (err CannotMoveCellError) Error() string {
	return fmt.Sprintf("Cannot move number %v to direction %v. number destination must be the empty spot.", err.number, err.direction)
}

type NumberNotExistError struct {
	number int
}

func (err NumberNotExistError) Error() string {
	return fmt.Sprintf("Number %v does not exist on the board", err.number)
}
