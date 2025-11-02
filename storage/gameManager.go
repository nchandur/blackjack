package storage

import (
	"fmt"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/users"
)

type GameManager struct {
	users    *[]users.User
	signedIn int
}

func NewGameManager(users *[]users.User) GameManager {
	gameManager := GameManager{users: users, signedIn: -1}

	for idx := range *users {
		if (*users)[idx].SignedIn {
			gameManager.signedIn = idx
		}
	}

	return gameManager
}

func (g GameManager) Save(game game.Game) error {

	if g.signedIn == -1 {
		return fmt.Errorf("please sign in")
	}

	(*g.users)[g.signedIn].Round.Played = game.Rounds
	(*g.users)[g.signedIn].Round.Won = game.Wins
	(*g.users)[g.signedIn].Round.Lost = game.Losses
	(*g.users)[g.signedIn].Round.Pushed = game.Pushes
	(*g.users)[g.signedIn].Kaasu = game.Kaasu

	return nil
}

func (g GameManager) Load() (game.Game, error) {

	toLoad := game.NewGame()

	if g.signedIn == -1 {
		return toLoad, fmt.Errorf("please sign in")
	}

	toLoad.Rounds = (*g.users)[g.signedIn].Round.Played
	toLoad.Wins = (*g.users)[g.signedIn].Round.Won
	toLoad.Losses = (*g.users)[g.signedIn].Round.Lost
	toLoad.Pushes = (*g.users)[g.signedIn].Round.Pushed
	toLoad.Kaasu = (*g.users)[g.signedIn].Kaasu

	return toLoad, nil
}
