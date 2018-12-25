package input

import (
	"../game"
	"../presentation"
	"fmt"
)

func ReadBoardSize() int {
	var boardSize int
	presentation.PrintChooseBoardSize()
	_, err := fmt.Scanf("%d", &boardSize)
	boardSize = ValidateBoardSize(boardSize, err)
	presentation.PrintSelectedBoardSize(boardSize)
	return boardSize
}

func ValidateBoardSize(num int, err error) int {
	if err != nil || num < game.MIN_BOARD_SIZE || num > game.MAX_BOARD_SIZE {
		return game.DEFAULT_BOARD_SIZE
	}

	return num
}

func ReadNumber(gameInstance game.Game) int {
	num := 0
	valid := false

	for !valid {
		presentation.PrintEnterNum(gameInstance)
		_, err := fmt.Scanf("%d", &num)
		num, valid = ValidateNum(num, err, gameInstance)
	}

	return num
}

func ValidateNum(num int, err error, gameInstance game.Game) (int, bool) {
	if err != nil {
		fmt.Println("This is not a valid number.", err)
	} else if num < 1 || num > gameInstance.NumbersCount() {
		fmt.Println(game.NumberNotExistError{num})
	} else {
		return num, true
	}

	return num, false
}

func ReadDirection() game.Direction {
	var direction game.Direction
	var dirAsText string
	valid := false

	for !valid {
		presentation.PrintEnterDirection()
		fmt.Scanf("%s\n", &dirAsText)
		direction, valid = ValidateDirection(dirAsText)
	}

	return direction
}

func ValidateDirection(dirAsText string) (game.Direction, bool) {
	direction := game.DirectionFromString(dirAsText)

	if direction == game.Direction("") {
		fmt.Println("Please enter a direction from the list")
		return direction, false
	}

	return direction, true

}
