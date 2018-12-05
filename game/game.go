package game

type Game struct {
	Board [][]int
}

func (g Game) String() string {
	return ""
}

func (myGame Game) MoveCell(number int, direction Direction) error {
	sourceX, sourceY, err := GetIndexesByNumber(myGame, number)

	if err != nil {
		return err
	}

	destX, destY, err := GetDestinationIndexes(myGame, sourceX, sourceY, direction)

	if err != nil {
		return err
	}

	moveCellInner(myGame, sourceX, sourceY, destX, destY)
	return nil
}

func moveCellInner(game Game, sourceX int, sourceY int, destX int, destY int) {
	game.Board[destX][destY] = game.Board[sourceX][sourceY]
	game.Board[sourceX][sourceY] = 0
}

func InitGame() Game {
	return Game{
		[][]int{},
	}
}

func (g Game) IsSuccess() bool {
	return false
}

func (g Game) NumbersCount() int {
	return len(g.Board)*len(g.Board[0]) - 1
}
