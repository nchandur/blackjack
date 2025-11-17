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

func (g GameManager) Save(game *game.Game) error {

	if g.signedIn == -1 {
		return fmt.Errorf("please sign in")
	}

	(*g.users)[g.signedIn].Stats.Played = game.Player.Played
	(*g.users)[g.signedIn].Stats.Won = game.Player.Won
	(*g.users)[g.signedIn].Stats.Lost = game.Player.Lost
	(*g.users)[g.signedIn].Stats.Pushed = game.Player.Pushed
	(*g.users)[g.signedIn].Kaasu = game.Player.Kaasu

	(*g.users)[0].Stats.Played += game.Dealer.Played
	(*g.users)[0].Stats.Won += game.Dealer.Won
	(*g.users)[0].Stats.Lost += game.Dealer.Lost
	(*g.users)[0].Stats.Pushed += game.Dealer.Pushed
	(*g.users)[0].Kaasu += game.Dealer.Kaasu

	return nil
}

func (g GameManager) Load() (*game.Game, error) {

	toLoad := game.NewGame()

	if g.signedIn == -1 {
		return toLoad, fmt.Errorf("please sign in")
	}

	toLoad.Player.Played = (*g.users)[g.signedIn].Stats.Played
	toLoad.Player.Won = (*g.users)[g.signedIn].Stats.Won
	toLoad.Player.Lost = (*g.users)[g.signedIn].Stats.Lost
	toLoad.Player.Pushed = (*g.users)[g.signedIn].Stats.Pushed
	toLoad.Player.Kaasu = (*g.users)[g.signedIn].Kaasu

	return toLoad, nil
}
