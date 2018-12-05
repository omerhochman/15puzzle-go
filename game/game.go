package game

type Game struct {
	Board [][]int
}

func (g Game) String() string {
	return ""
}

func (g Game) MoveCell(number int, direction Direction) error {
	return nil
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
