package player

import (
	"github.com/nchandur/blackjack/game/structs"
)

type Player struct {
	structs.Hand
}

func NewPlayer(shoe *structs.Shoe) Player {
	return Player{Hand: structs.NewHand(shoe)}
}
