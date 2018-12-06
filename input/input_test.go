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

	assertNumberInvalid(myGame, t, "0")
	assertNumberInvalid(myGame, t, "9")
}

func assertNumberInvalid(myGame game.Game, t *testing.T, numberAsText string) {
	_, valid := ValidateNum(numberAsText, myGame)
	if valid {
		t.Error("Number should be invalid")
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

	assertNumberValid(myGame, t, "1", 1)
	assertNumberValid(myGame, t, "8", 8)
}

func assertNumberValid(myGame game.Game, t *testing.T, numberAsText string, expectedNumber int) {
	num, valid := ValidateNum(numberAsText, myGame)
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
