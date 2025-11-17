package dealer

import "github.com/nchandur/blackjack/model"

type Dealer struct {
	model.Hand
	model.Stats
	Kaasu int64
}

func NewDealer(shoe *model.Shoe) *Dealer {
	dealer := Dealer{Hand: model.NewHand(shoe)}
	return &dealer
}
