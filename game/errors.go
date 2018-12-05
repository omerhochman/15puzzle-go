package game

import "fmt"

type CannotMoveCellError struct {
	number    int
	direction Direction
}

func (err CannotMoveCellError) Error() string {
	return fmt.Sprint("Cannot move number %v to direction %v. number destination must be the empty spot.", err.number, err.direction)
}

type NumberNotExistError struct {
	number int
}

func (err NumberNotExistError) Error() string {
	return fmt.Sprint("Number %v does not exist on the board", err.number)
}
