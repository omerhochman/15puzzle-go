package input

import (
	"../game"
	"fmt"
)

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
