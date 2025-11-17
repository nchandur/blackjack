package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nchandur/blackjack/cmd"
	"github.com/nchandur/blackjack/storage"
	"github.com/nchandur/blackjack/users"
)

func main() {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	cmd.GAME_DIR = filepath.Join(homeDir, "blackjack")

	if _, err := os.Stat(cmd.GAME_DIR); os.IsNotExist(err) {
		err := os.Mkdir(cmd.GAME_DIR, 0777)

		if err != nil {
			log.Fatal(err)
		}

	}

	uses := make([]users.User, 0)

	store := storage.NewStorage(filepath.Join(cmd.GAME_DIR, "save.json"))

	err = store.Load(&uses)

	if err != nil {
		log.Fatal(err)
	}

	um := storage.NewUserManager(&uses)
	gm := storage.NewGameManager(&uses)

	cmd.Execute(um, gm)

	store.Save(uses)
}
