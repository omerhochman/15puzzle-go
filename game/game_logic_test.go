package game

import "testing"

func assertIndexes(expectedI int, expectedJ int, actualI int, actualJ int, t *testing.T) {
	if actualI != expectedI || expectedJ != actualJ {
		t.Errorf("Expected indexes %v, %v and recieved: %v, %v", expectedI, expectedJ, actualI, actualJ)
	}
}

func TestGetDestinationIndexes(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	game := Game{
		board,
	}

	i, j, err := GetDestinationIndexes(game, 0, 1, DOWN)

	if err != nil {
		t.Errorf("Expected null error but recieved %v", err)
	}

	assertIndexes(1, 1, i, j, t)
}

func TestGetDestinationIndexes_Failure(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	game := Game{
		board,
	}

	_, _, err := GetDestinationIndexes(game, 0, 1, UP)

	if _, ok := err.(CannotMoveCellError); !ok {
		t.Errorf("Expected CannotMoveCellError but recieved %v", err)
	}
}

func TestGetIndexesByNumber_Success(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	game := Game{
		board,
	}
	assertGetIndexesSuccess(game, t, 5, 2, 1)
	assertGetIndexesSuccess(game, t, 1, 0, 0)
	assertGetIndexesSuccess(game, t, 7, 2, 0)

}

func assertGetIndexesSuccess(game Game, t *testing.T, number, expectedI, expectedJ int) {
	i, j, err := GetIndexesByNumber(game, number)
	if err != nil {
		t.Errorf("Expected null error but recieved %v", err)
	}
	assertIndexes(expectedI, expectedJ, i, j, t)
}

func TestGetIndexesByNumber_Failure(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}

	game := Game{
		board,
	}
	assertGetIndexesFailure(game, t, 9)
	assertGetIndexesFailure(game, t, 10)
	assertGetIndexesFailure(game, t, 0)

}

func assertGetIndexesFailure(game Game, t *testing.T, number int) {
	_, _, err := GetIndexesByNumber(game, number)
	if _, ok := err.(NumberNotExistError); !ok {
		t.Errorf("Expected NumberNotExistError but recieved %v", err)
	}
}
