package validations

import game "../game"

func (myGame game.Game) CanMoveCell(number int, direction game.Direction) (bool, error) {
	return false, nil
}

func getIndexesByNumber(myGame game.Game, number int) (i int, j int, err error) {
	for i := range myGame.Board {
		for j := range myGame.Board[i] {
			if myGame.Board[i][j] == number {
				return i, j, nil
			}
		}
	}

	return 0, 0, game.NumberNotExistError{number}
}
