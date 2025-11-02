package dealer

import "github.com/nchandur/blackjack/game/structs"

func (d *Dealer) Hit(shoe *structs.Shoe) {
	d.Hand = append(d.Hand, shoe.Draw())
	d.Hand.MakeOptimum()
}
