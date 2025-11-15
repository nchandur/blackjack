package bot

import (
	"github.com/nchandur/blackjack/model"
)

func (b *Bot) Play(bet int64, dealer uint8, shoe *model.Shoe) int64 {

	stop := false

	for !stop {

		if b.IsBlackjack() && len(b.Hand) == 2 {
			return bet + (bet / 2)
		}

		if b.IsBlackjack() {
			return bet
		}

		if b.IsBust() {
			return bet
		}

		action := b.Lookup(dealer)

		switch action {
		case 0:
			stop = true
		case 1:
			b.Hit(shoe)
		case 2:
			if len(b.Hand) == 2 && b.Kaasu >= (2*bet) {
				bet *= 2
				b.Hit(shoe)
			}
			stop = true
		}

	}

	return bet

}
