package player

import (
	"github.com/nchandur/blackjack/entities"
	"github.com/nchandur/blackjack/game/structs"
)

type Player struct {
	structs.Hand
	Kaasu int
	entities.Stats
}

func NewPlayer(shoe *structs.Shoe) *Player {
	player := Player{
		Hand:  structs.NewHand(shoe),
		Kaasu: 1000,
	}

	return &player
}
