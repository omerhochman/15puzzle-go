package presentation

import (
	"../game"
	"fmt"
)

func PrintMenu(myGame game.Game, err error) string {
	prefix := myGame.String()

	if err != nil {
		prefix = err.Error()
	}

	text := fmt.Sprintf("\n%v\nPlease enter a number and then a direction.\n", prefix)
	fmt.Print(text)
	return text
}

func PrintSuccess(myGame game.Game) {
	fmt.Printf("\n%v\nCongratulations!", myGame)
}
