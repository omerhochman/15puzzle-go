package presentation

import (
	"../game"
	"bytes"
	"fmt"
)

const MENU_FORMAT = "\n\n%v\nPlease enter a number and then a direction.\n"

func PrintMenu(myGame game.Game, err error) string {
	var buffer bytes.Buffer
	buffer.WriteString(myGame.String())

	if err != nil {
		buffer.WriteString(err.Error())
	}

	text := fmt.Sprintf(MENU_FORMAT, buffer.String())
	fmt.Print(text)
	return text
}

func PrintSuccess(myGame game.Game) {
	fmt.Printf("\n%v\nCongratulations!", myGame)
}
