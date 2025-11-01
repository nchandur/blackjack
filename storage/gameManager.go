package storage

import (
	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/player"
)

type GameManager struct {
	players  *[]player.Player
	signedIn int
	gameID   int
}

func NewGameManager(players *[]player.Player) GameManager {
	gameManager := GameManager{players: players}

	for idx := range *players {
		if (*players)[idx].SignedIn {
			gameManager.signedIn = idx
		}
	}

	return gameManager
}

func (g GameManager) Save(game game.Game) error {
	return nil
}

func (g GameManager) Load(gameID int) (game.Game, error) {
	return game.Game{}, nil
}
