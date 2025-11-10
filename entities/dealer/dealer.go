package dealer

import (
	"github.com/nchandur/blackjack/entities"
	"github.com/nchandur/blackjack/game/structs"
)

type Dealer struct {
	structs.Hand
	Kaasu int
	entities.Stats
}

func NewDealer(shoe *structs.Shoe) *Dealer {
	dealer := Dealer{
		Hand: structs.NewHand(shoe),
	}

	return &dealer
}
