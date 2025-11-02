package main

import (
	"log"

	"github.com/nchandur/blackjack/cmd"
	"github.com/nchandur/blackjack/storage"
	"github.com/nchandur/blackjack/users"
)

func main() {

	uses := make([]users.User, 0)

	store := storage.NewStorage("blackjack.json")

	err := store.Load(&uses)

	if err != nil {
		log.Fatal(err)
	}

	um := storage.NewUserManager(&uses)
	gm := storage.NewGameManager(&uses)

	cmd.Execute(um, gm)

	store.Save(uses)
}
