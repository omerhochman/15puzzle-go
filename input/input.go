package input

import (
	"../game"
	"fmt"
)

func ReadBoardSize() int {
	var boardSize int
	fmt.Printf("Please choose board size (%v-%v, or anything else for default size %v): ", game.MIN_BOARD_SIZE, game.MAX_BOARD_SIZE, game.DEFAULT_BOARD_SIZE)
	_, err := fmt.Scanf("%d", &boardSize)
	boardSize = ValidateBoardSize(boardSize, err)
	fmt.Println("Playing with board size: ", boardSize)
	return boardSize
}

func ValidateBoardSize(num int, err error) int {
	if err != nil || num < game.MIN_BOARD_SIZE || num > game.MAX_BOARD_SIZE {
		return game.DEFAULT_BOARD_SIZE
	}

	return num
}

func ReadNumber(myGame game.Game) int {
	num := 0
	valid := false

	for !valid {
		fmt.Printf("Enter a number (1-%v): ", myGame.NumbersCount())
		_, err := fmt.Scanf("%d", &num)
		num, valid = ValidateNum(num, err, myGame)
	}

	return num
}

func ValidateNum(num int, err error, myGame game.Game) (int, bool) {
	if err != nil {
		fmt.Println("This is not a valid number.", err)
	} else if num < 1 || num > myGame.NumbersCount() {
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
		fmt.Print("Enter a direction (up, down, left or right): ")

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
