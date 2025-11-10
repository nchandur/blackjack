package game

import (
	"github.com/nchandur/blackjack/entities/dealer"
	"github.com/nchandur/blackjack/entities/player"
	"github.com/nchandur/blackjack/game/structs"
)

var SHOE_SIZE = 6

type Game struct {
	Player *player.Player
	Dealer *dealer.Dealer
	shoe   *structs.Shoe
}

func NewGame() Game {
	shoe := structs.NewShoe(6)

	play := player.NewPlayer(&shoe)
	deal := dealer.NewDealer(&shoe)

	return Game{
		Player: play,
		Dealer: deal,
		shoe:   &shoe,
	}
}

