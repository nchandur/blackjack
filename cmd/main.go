package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/storage"
	"github.com/nchandur/blackjack/users"
)

var um users.UserManager
var gm game.GameManager

func main() {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	uses := make([]users.User, 0)

	store := storage.NewStorage(filepath.Join(homeDir, ".blackjack.save"))

	err = store.Load(&uses)

	if err != nil {
		log.Fatal(err)
	}

	um = storage.NewUserManager(&uses)
	gm = storage.NewGameManager(&uses)

	Execute()

	store.Save(uses)
}
