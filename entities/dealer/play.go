package dealer

import "github.com/nchandur/blackjack/game/structs"

func (d *Dealer) Play(shoe *structs.Shoe, bet int) int {

	for d.FindSum() <= 17 {
		if d.CheckBlackjack() {
			break
		}

		if d.CheckBust() {
			break
		}

		d.Hit(shoe)

	}

	return bet

}
