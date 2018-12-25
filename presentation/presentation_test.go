package presentation

import (
	"../game"
	"fmt"
	"strings"
	"testing"
)

func TestGame_String(t *testing.T) {
	gameInstance := game.Game{[][]int{
		{1, 2, 3},
		{4, 15, 6},
		{7, 8, 0},
	}}

	result := gameInstance.String()
	expectedResult := "   1    2    3\n" +
		"   4   15    6\n" +
		"   7    8    _\n"
	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Printed board is not as expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)

	}
}

func TestPrintMenu_no_error(t *testing.T) {
	gameInstance := game.Game{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}}

	result := PrintMenu(gameInstance, nil)
	expectedResult := fmt.Sprintf(MENU_FORMAT, gameInstance)

	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Printed menu is not as expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)
	}
}

func TestPrintMenu_error(t *testing.T) {
	gameInstance := game.Game{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}}

	cellError := game.NumberNotExistError{}
	result := PrintMenu(gameInstance, cellError)

	expectedResult := fmt.Sprintf(MENU_FORMAT, gameInstance.String()+"\n"+cellError.Error())

	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Error expected. Actual:\n%v\nExpected:\n%v", result, expectedResult)
	}
}
