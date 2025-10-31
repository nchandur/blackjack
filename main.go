package main

import (
	"log"

	"github.com/nchandur/blackjack/cmd"
	"github.com/nchandur/blackjack/player"
	"github.com/nchandur/blackjack/storage"
)

func main() {

	players := make([]player.Player, 0)

	store := storage.NewStorage("blackjack.json")

	err := store.Load(&players)

	if err != nil {
		log.Fatal(err)
	}

	playerManager := storage.NewPlayerManager(&players)

	cmd.Execute(playerManager)

	store.Save(players)

}
