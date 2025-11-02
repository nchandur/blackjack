package player

import "github.com/nchandur/blackjack/game/structs"

func (p *Player) Hit(shoe *structs.Shoe) {
	p.Hand = append(p.Hand, shoe.Draw())
	p.Hand.MakeOptimum()
}
