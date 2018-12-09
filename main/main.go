package main

import (
	"../game"
	"../input"
	"../presentation"
)

func main() {
	gameFlow()
}

func gameFlow() {
	presentation.PrintWelcome()
	boardSize := input.ReadBoardSize()
	myGame := game.InitGame(boardSize)
	var err error
	for !myGame.IsSuccess() {
		presentation.PrintMenu(myGame, err)
		number := input.ReadNumber(myGame)
		direction := input.ReadDirection()
		err = myGame.MoveCell(number, direction)
	}
	presentation.PrintSuccess(myGame)
}
