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
	myGame := game.InitGame()
	var err error
	for !myGame.IsSuccess() {
		presentation.PrintMenu(myGame, err)
		number := input.ReadNumber()
		direction := input.ReadDirection()
		err = myGame.MoveCell(number, direction)
	}
	presentation.PrintSuccess()
}
