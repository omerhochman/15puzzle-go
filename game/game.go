package game

import (
	"bytes"
	"fmt"
)

type Game struct {
	Board [][]int
}

func (g Game) String() string {
	var buffer bytes.Buffer

	for _, row := range g.Board {
		for j, cell := range row {
			if cell != 0 {
				buffer.WriteString(fmt.Sprintf("%4d", cell))
			} else {
				buffer.WriteString(fmt.Sprintf("%4v", "_"))
			}

			if j < len(row)-1 {
				buffer.WriteString(" ")
			}
		}

		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (gameInstance Game) MoveCell(number int, direction Direction) error {
	sourceRow, sourceCol, err := GetIndexesByNumber(gameInstance, number)

	if err != nil {
		return err
	}

	destRow, destCol, err := GetDestinationIndexes(gameInstance, sourceRow, sourceCol, direction, number)

	if err != nil {
		return err
	}

	gameInstance.Board[destRow][destCol] = gameInstance.Board[sourceRow][sourceCol]
	gameInstance.Board[sourceRow][sourceCol] = 0

	return nil
}

func InitGame(boardSize int) Game {
	var gameInstance Game

	for {
		board := CreateRandomBoard(boardSize)

		gameInstance = Game{
			board,
		}

		if !gameInstance.IsSuccess() {
			break
		}
	}

	return gameInstance
}

func (g Game) IsSuccess() bool {
	num := 1
	for _, row := range g.Board {
		for _, cell := range row {
			if cell != 0 && cell != num {
				return false
			}
			num++
		}
	}

	return true
}

func (g Game) NumbersCount() int {
	return len(g.Board)*len(g.Board[0]) - 1
}
