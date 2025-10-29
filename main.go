package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/nchandur/blackjack/models"
)

func main() {

	message := "Welcome to Blackjack!!!"

	repeated := strings.Repeat("-", utf8.RuneCountInString(message))

	fmt.Println(message)
	fmt.Println(repeated)

	game := models.NewGame()
	game.PlayRound()

}
