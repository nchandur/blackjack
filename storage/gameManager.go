package storage

import (
	"fmt"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/player"
)

type GameManager struct {
	players  *[]player.Player
	signedIn int
}

func NewGameManager(players *[]player.Player) GameManager {
	gameManager := GameManager{players: players, signedIn: -1}

	for idx := range *players {
		if (*players)[idx].SignedIn {
			gameManager.signedIn = idx
		}
	}

	return gameManager
}

func (g GameManager) Save(game game.Game) error {

	if g.signedIn == -1 {
		return fmt.Errorf("please sign in")
	}

	(*g.players)[g.signedIn].Round.Played = game.Rounds
	(*g.players)[g.signedIn].Round.Won = game.Wins
	(*g.players)[g.signedIn].Round.Lost = game.Losses
	(*g.players)[g.signedIn].Round.Pushed = game.Pushes
	(*g.players)[g.signedIn].Buckaroonies = game.Buckaroonies

	return nil
}

func (g GameManager) Load() (game.Game, error) {

	toLoad := game.NewGame()

	if g.signedIn == -1 {
		return toLoad, fmt.Errorf("please sign in")
	}

	toLoad.Rounds = (*g.players)[g.signedIn].Round.Played
	toLoad.Wins = (*g.players)[g.signedIn].Round.Won
	toLoad.Losses = (*g.players)[g.signedIn].Round.Lost
	toLoad.Pushes = (*g.players)[g.signedIn].Round.Pushed
	toLoad.Buckaroonies = (*g.players)[g.signedIn].Buckaroonies

	return toLoad, nil
}
