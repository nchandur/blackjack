package main

import (
	"log"

	"github.com/nchandur/blackjack/game"
)

func main() {

	game := game.NewGame()

	if err := game.Play(); err != nil {
		log.Fatal(err)
	}

}
