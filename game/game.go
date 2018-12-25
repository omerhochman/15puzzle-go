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

func (myGame Game) MoveCell(number int, direction Direction) error {
	sourceRow, sourceCol, err := GetIndexesByNumber(myGame, number)

	if err != nil {
		return err
	}

	destX, destY, err := GetDestinationIndexes(myGame, sourceRow, sourceCol, direction, number)

	if err != nil {
		return err
	}

	moveCellInner(myGame, sourceRow, sourceCol, destX, destY)

	return nil
}

func moveCellInner(game Game, sourceX int, sourceY int, destX int, destY int) {
	game.Board[destX][destY] = game.Board[sourceX][sourceY]
	game.Board[sourceX][sourceY] = 0
}

func InitGame(boardSize int) Game {
	var myGame Game

	for {
		board := CreateRandomBoard(boardSize)

		myGame = Game{
			board,
		}

		if !myGame.IsSuccess() {
			break
		}
	}

	return myGame
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
