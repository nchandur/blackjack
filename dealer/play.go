package dealer

import "github.com/nchandur/blackjack/model"

func (d *Dealer) Play(bet int64, shoe *model.Shoe) int64 {
	for d.Hand.Sum() < 17 {
		d.Hit(shoe)
	}

	return bet
}
