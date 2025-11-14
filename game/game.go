package game

import (
	"github.com/nchandur/blackjack/dealer"
	"github.com/nchandur/blackjack/model"
	"github.com/nchandur/blackjack/player"
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
