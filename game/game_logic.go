package game

func ValidateDestinationIndexes(myGame Game, destX int, destY int) error {
	return nil
}

func GetDestinationIndexes(myGame Game, sourceX int, sourceY int, direction Direction) (i int, j int, err error) {
	return 0, 0, nil
}

func GetIndexesByNumber(myGame Game, number int) (i int, j int, err error) {
	for i := range myGame.Board {
		for j := range myGame.Board[i] {
			if myGame.Board[i][j] == number {
				return i, j, nil
			}
		}
	}

	return 0, 0, NumberNotExistError{number}
}
