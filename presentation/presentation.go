package presentation

import (
	"../game"
	"bytes"
	"fmt"
)

const WELCOME = "Welcome to the '15 puzzle' game, and good luck!"
const CHOOSE_BOARD_SIZE = "Please choose board size (%v-%v, or anything else for default size %v): "
const PLEASE_WAIT = "\nInitializing your board...\n"
const SELECTED_BOARD_SIZE = "Playing with board size: "
const MENU_FORMAT = "\n\n%v\nPlease enter a number and then a direction.\n"
const ENTER_NUM = "Enter a number (1-%v): "
const ENTER_DIRECTION = "Enter a direction (up, down, left or right): "
const CONGRATULATIONS = "\n%v\nCongratulations!\n\n"

func PrintMenu(gameInstance game.Game, err error) string {
	var buffer bytes.Buffer
	buffer.WriteString(gameInstance.String())

	if err != nil {
		buffer.WriteString("\n")
		buffer.WriteString(err.Error())
	}

	text := fmt.Sprintf(MENU_FORMAT, buffer.String())
	fmt.Print(text)
	return text
}

func PrintPleaseWait() {
	fmt.Println(PLEASE_WAIT)
}

func PrintWelcome() {
	fmt.Println(WELCOME)
}

func PrintSuccess(gameInstance game.Game) {
	fmt.Printf(CONGRATULATIONS, gameInstance)
}

func PrintChooseBoardSize() {
	fmt.Printf(CHOOSE_BOARD_SIZE, game.MIN_BOARD_SIZE, game.MAX_BOARD_SIZE, game.DEFAULT_BOARD_SIZE)

}

func PrintSelectedBoardSize(boardSize int) {
	fmt.Println(SELECTED_BOARD_SIZE, boardSize)
}

func PrintEnterNum(gameInstance game.Game) {
	fmt.Printf(ENTER_NUM, gameInstance.NumbersCount())
}

func PrintEnterDirection() {
	fmt.Print(ENTER_DIRECTION)
}
