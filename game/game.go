package game

import (
	"github.com/nchandur/blackjack/entities/dealer"
	"github.com/nchandur/blackjack/entities/player"
	"github.com/nchandur/blackjack/model"
)

type Game struct {
	player.Player
	dealer.Dealer
	shoe model.Shoe
}

func NewGame() *Game {
	game := Game{
		shoe: model.NewShoe(),
	}

	game.Player = *player.NewPlayer(&game.shoe)
	game.Dealer = *dealer.NewDealer(&game.shoe)

	return &game
}

type GameManager interface {
	Save(game *Game) error
	Load() (*Game, error)
}
