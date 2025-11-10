package player

import (
	"fmt"

	"github.com/nchandur/blackjack/game/structs"
)

func (p *Player) Play(shoe *structs.Shoe, bet int) int {

	stop := false

	for ok := true; ok; ok = !stop {

		fmt.Printf("\nPlayer Hand: %d\n%s", p.Hand.FindSum(), p.Hand.String())

		if len(p.Hand) == 2 && p.CheckBlackjack() {
			bet = (bet) + (bet / 2)
			return bet
		}

		if p.CheckBlackjack() {
			break
		}

		if p.CheckBust() {
			break
		}

		fmt.Printf("Your move (h/s/d/q): ")

		ip := ""

		_, err := fmt.Scanln(&ip)

		if err != nil {
			return 0
		}

		switch ip {
		case "h", "hit":
			p.Hit(shoe)
		case "s", "stand":
			stop = true
		case "d", "double":
			if (len(p.Hand) == 2) && (2*bet < p.Kaasu) {
				bet = 2 * bet
				p.Hit(shoe)
				stop = true
			}
		case "q", "quit":
			return -1
		default:
			fmt.Println("invalid input")
		}

	}

	return bet
}
