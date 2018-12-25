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
	presentation.PrintPleaseWait()
	gameInstance := game.InitGame(boardSize)
	var err error
	for !gameInstance.IsSuccess() {
		presentation.PrintMenu(gameInstance, err)
		number := input.ReadNumber(gameInstance)
		direction := input.ReadDirection()
		err = gameInstance.MoveCell(number, direction)
	}
	presentation.PrintSuccess(gameInstance)
}
