package game

import (
	"reflect"
	"testing"
)

func findNumInBoard(board [][]int, number int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == number {
				return true
			}
		}
	}

	return false
}

func TestInitGame(t *testing.T) {
	game := InitGame()
	for x := 0; x < len(game.board)*len(game.board[0]); x++ {
		if !findNumInBoard(game.board, x) {
			t.Errorf("Number %v wasnt found in board", x)
		}
	}
}

func TestGame_IsSuccess_success(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	game := Game{
		board,
	}

	if !game.IsSuccess() {
		t.Errorf("Game should be succeeded")

	}
}

func TestGame_IsSuccess_failure(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	}

	game := Game{
		board,
	}

	if game.IsSuccess() {
		t.Errorf("Game should be failed")
	}
}

func TestGame_MoveCell_succeeded(t *testing.T) {

	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	game := Game{
		board,
	}

	assertMoveSuccessfully(
		5, UP, game, t, [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 0, 8},
		})
	assertMoveSuccessfully(
		2, DOWN, game, t, [][]int{
			{1, 0, 3},
			{4, 2, 6},
			{7, 5, 8},
		})
	assertMoveSuccessfully(
		6, LEFT, game, t, [][]int{
			{1, 2, 3},
			{4, 6, 0},
			{7, 5, 8},
		})
	assertMoveSuccessfully(
		4, RIGHT, game, t, [][]int{
			{1, 2, 3},
			{0, 4, 6},
			{7, 5, 8},
		})
}

func assertMoveSuccessfully(number int, dir Direction, game Game, t *testing.T, expected_board [][]int) {
	err := game.MoveCell(number, dir)

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected_board, game.board) {
		t.Errorf("Board before moving cell: %v, expected after moving cell: %v", game.board, expected_board)
	}
}

func TestGame_MoveCell_failures(t *testing.T) {

	board := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	}

	game := Game{
		board,
	}

	assertMoveFailed(game, 8, DOWN, t)
	assertMoveFailed(game, 8, UP, t)
	assertMoveFailed(game, 0, UP, t)
	assertMoveFailed(game, 2, UP, t)
	assertMoveFailed(game, 1, LEFT, t)
	assertMoveFailed(game, 3, RIGHT, t)
}

func assertMoveFailed(game Game, i int, direction Direction, t *testing.T) {
	err := game.MoveCell(i, direction)
	if _, ok := err.(CannotMoveCellError); !ok {
		t.Errorf("There should be an error of CannotMoveCellError")
	}
}
