package game

func GetDestinationIndexes(myGame Game, sourceRow int, sourceCol int, direction Direction, number int) (int, int, error) {
	row, col := getDestinationIndexesInner(direction, sourceRow, sourceCol)

	if !areDestinationIndexesValid(myGame, row, col) {
		return row, col, CannotMoveCellError{number, direction}
	}

	return row, col, nil
}

func areDestinationIndexesValid(myGame Game, destRow int, destCol int) bool {
	if destRow >= len(myGame.Board) || destRow < 0 {
		return false
	}

	if destCol >= len(myGame.Board[0]) || destCol < 0 {
		return false
	}

	if myGame.Board[destRow][destCol] != 0 {
		return false
	}

	return true
}

func getDestinationIndexesInner(direction Direction, sourceRow int, sourceCol int) (int, int) {
	row, col := 0, 0
	switch direction {
	case UP:
		row, col = sourceRow-1, sourceCol
	case DOWN:
		row, col = sourceRow+1, sourceCol
	case LEFT:
		row, col = sourceRow, sourceCol-1
	case RIGHT:
		row, col = sourceRow, sourceCol+1
	}
	return row, col
}

func GetIndexesByNumber(myGame Game, number int) (i int, j int, err error) {
	if number == 0 {
		return 0, 0, NumberNotExistError{0}
	}

	for i, row := range myGame.Board {
		for j := range row {
			if myGame.Board[i][j] == number {
				return i, j, nil
			}
		}
	}

	return 0, 0, NumberNotExistError{number}
}
