package player

import (
	"fmt"

	"github.com/nchandur/blackjack/model"
)

func (p *Player) Play(bet int64, shoe *model.Shoe) int64 {

	if bet > p.Kaasu {
		fmt.Println("insufficient funds")
		return 0
	}

	stop := false

	for !stop {

		fmt.Printf("Your hand: \n%s\nSum: %d\n", p.Hand.String(), p.Hand.Sum())

		if p.IsBlackjack() && len(p.Hand) == 2 {
			return bet + (bet / 2)
		}

		if p.IsBlackjack() {
			return bet
		}

		if p.IsBust() {
			return bet
		}

		var input string
		fmt.Printf("Your move (h/s/d/q): ")
		fmt.Scanln(&input)

		switch input {
		case "h", "hit":
			p.Hit(shoe)
		case "d", "double":
			if len(p.Hand) == 2 && p.Kaasu >= (2*bet) {
				bet *= 2
				p.Hit(shoe)
			}
			stop = true
		case "s", "stand":
			stop = true
		}

		if input == "q" {
			return 0
		}

	}

	return bet

}
