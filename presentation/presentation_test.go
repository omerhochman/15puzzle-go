package presentation

import (
	"../game"
	"fmt"
	"strings"
	"testing"
)

func TestGame_String(t *testing.T) {
	myGame := game.Game{[][]int{
		{1, 2, 3},
		{4, 15, 6},
		{7, 8, 0},
	}}

	result := myGame.String()
	expectedResult := "   1    2    3\n" +
		"   4   15    6\n" +
		"   7    8    _\n"
	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Printed board is not as expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)

	}
}

func TestPrintMenu_no_error(t *testing.T) {
	myGame := game.Game{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}}

	result := PrintMenu(myGame, nil)
	expectedResult := fmt.Sprintf(
		myGame.String()+
			"\nPlease enter a number(1-%v) and then a direction (up, down, left, right)\n", myGame.NumbersCount())

	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Printed menu is not as expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)
	}
}

func TestPrintMenu_error(t *testing.T) {
	myGame := game.Game{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}}

	cellError := game.NumberNotExistError{}
	result := PrintMenu(myGame, cellError)

	expectedResult := fmt.Sprintf(
		"%v\nPlease enter a number(1-%v) and then a direction (up, down, left, right)\n", cellError, myGame.NumbersCount())

	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Error expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)
	}
}
