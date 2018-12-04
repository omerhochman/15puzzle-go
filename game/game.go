package game

const GAME_WIDTH = 4
const GAME_HEIGHT = 4

type Game struct {
	board []int
}

func (g Game) String() string {
	return ""
}

func (g Game) moveCell(number int, direction string) error {
	return nil
}

func (g Game) init() {
}

func (g Game) isSuccess() bool {
	return false
}
