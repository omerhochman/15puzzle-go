package game

type Game struct {
	board [][]int
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
