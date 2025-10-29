package main

import (
	"github.com/nchandur/blackjack/models"
)

func main() {

	game := models.NewGame()
	game.PlayRound()

}
