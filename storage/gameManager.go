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

	(*g.users)[g.signedIn].Stats.Played = game.Stats.Rounds
	(*g.users)[g.signedIn].Stats.Wins = game.Stats.Wins
	(*g.users)[g.signedIn].Stats.Losses = game.Stats.Losses
	(*g.users)[g.signedIn].Stats.Pushes = game.Stats.Pushes
	(*g.users)[g.signedIn].Kaasu = game.Kaasu

	return nil
}

func (g GameManager) Load() (game.Game, error) {

	toLoad := game.NewGame()

	if g.signedIn == -1 {
		return toLoad, fmt.Errorf("please sign in")
	}

	toLoad.Stats.Rounds = (*g.users)[g.signedIn].Stats.Played
	toLoad.Stats.Wins = (*g.users)[g.signedIn].Stats.Wins
	toLoad.Stats.Losses = (*g.users)[g.signedIn].Stats.Losses
	toLoad.Stats.Pushes = (*g.users)[g.signedIn].Stats.Pushes
	toLoad.Kaasu = (*g.users)[g.signedIn].Kaasu

	return toLoad, nil
}
