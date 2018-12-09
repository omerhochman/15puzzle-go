package input

import (
	"../game"
	"testing"
)

func TestValidateNum_invalid(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	myGame := game.Game{
		board,
	}

	assertNumberInvalid(myGame, t, 0)
	assertNumberInvalid(myGame, t, 9)
}

func assertNumberInvalid(myGame game.Game, t *testing.T, number int) {
	_, valid := ValidateNum(number, nil, myGame)
	if valid {
		t.Error("Number should be invalid")
	}
}

func TestValidateNum_input_error(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	myGame := game.Game{
		board,
	}

	_, valid := ValidateNum(1, game.NumberNotExistError{}, myGame)

	if valid {
		t.Error("There was an error in the reading process, number should be invalid")
	}
}

func TestValidateNum_valid(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	myGame := game.Game{
		board,
	}

	assertNumberValid(myGame, t, 1)
	assertNumberValid(myGame, t, 8)
}

func assertNumberValid(myGame game.Game, t *testing.T, expectedNumber int) {
	num, valid := ValidateNum(expectedNumber, nil, myGame)
	if !valid {
		t.Error("Number should be valid")
	}
	if num != expectedNumber {
		t.Errorf("Number is not as expected. Actual: %v. Expected: %v.", num, expectedNumber)
	}
}

func TestValidateDirection_invalid(t *testing.T) {
	assertDirectionInvalid(t, "lefty")
}

func assertDirectionInvalid(t *testing.T, directionAsText string) {
	_, valid := ValidateDirection(directionAsText)
	if valid {
		t.Error("Number should be invalid")
	}
}

func TestValidateDirection_valid(t *testing.T) {
	assertDirectionValid(t, "up", game.UP)
}

func assertDirectionValid(t *testing.T, directionAsText string, expectedDirection game.Direction) {
	num, valid := ValidateDirection(directionAsText)
	if !valid {
		t.Error("Direction should be valid")
	}
	if num != expectedDirection {
		t.Errorf("Direction is not as expected. Actual: %v. Expected: %v.", num, expectedDirection)
	}
}

func TestValidateBoardSize(t *testing.T) {
	assertBoardSize(1, game.DEFAULT_BOARD_SIZE, t)
	assertBoardSize(11, game.DEFAULT_BOARD_SIZE, t)
	assertBoardSize(2, 2, t)
	assertBoardSize(5, 5, t)
}

func assertBoardSize(userInput, expectedSize int, t *testing.T) {
	boardSize := ValidateBoardSize(userInput, nil)
	if boardSize != expectedSize {
		t.Errorf("Expected board size is: %v", expectedSize)
	}
}
