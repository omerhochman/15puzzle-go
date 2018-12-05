package game

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Board [][]int
}

func (g Game) String() string {
	var buffer bytes.Buffer

	for i, row := range g.Board {
		for j, cell := range row {
			buffer.WriteString(fmt.Sprintf("%4d", cell))

			if j < len(row)-1 {
				buffer.WriteString(" ")
			}
		}
		if i < len(g.Board)-1 {
			buffer.WriteString("\n")
		}
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

func InitGame() Game {
	numbersArray, k := createShuffledNumbersArray(), 0

	board := make([][]int, GAME_HEIGHT)

	for i := range board {
		board[i] = make([]int, GAME_WIDTH)
		for j := range board[i] {
			board[i][j] = numbersArray[k]
			k++
		}
	}

	return Game{
		board,
	}
}

func createShuffledNumbersArray() []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	arr := make([]int, GAME_HEIGHT*GAME_WIDTH)

	for i, num := range r.Perm(len(arr)) {
		arr[i] = num
		num++
	}

	return arr
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
