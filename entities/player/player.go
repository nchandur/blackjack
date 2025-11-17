package player

import (
	"github.com/nchandur/blackjack/model"
)

type Player struct {
	model.Hand
	model.Stats
	Kaasu int64
}

func NewPlayer(shoe *model.Shoe) *Player {
	player := Player{Hand: model.NewHand(shoe), Kaasu: 1000}
	return &player
}
