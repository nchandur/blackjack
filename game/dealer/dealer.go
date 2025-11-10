package dealer

import "github.com/nchandur/blackjack/game/structs"

type Dealer struct {
	structs.Hand
}

func NewDealer(shoe *structs.Shoe) Dealer {
	return Dealer{Hand: structs.NewHand(shoe)}
}
