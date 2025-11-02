package player

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/nchandur/blackjack/game/structs"
)

func (p *Player) Play(reader *bufio.Reader, shoe *structs.Shoe, bet, funds int) (string, int) {

	stop := false

	for ok := true; ok; ok = !stop {

		fmt.Printf("Player Hand\n%sSum: %d\n", p.Hand.String(), p.Hand.FindSum())

		if len(p.Hand) == 2 && p.CheckBlackjack() {
			bet = (bet) + (bet / 2)
		}

		if p.CheckBlackjack() {
			break
		}

		if p.CheckBust() {
			break
		}

		fmt.Printf("Your move (h/s/d/q): ")
		input, _ := reader.ReadString('\n')

		input = strings.ToLower(strings.Trim(input, "\n"))

		switch input {
		case "h", "hit":
			p.Hit(shoe)
		case "s", "stand":
			stop = true
		case "d", "double":
			if (len(p.Hand) == 2) && (2*bet < funds) {
				bet = 2 * bet
				p.Hit(shoe)
				stop = true
			}
		case "q", "quit":
			return "q", bet
		default:
			fmt.Println("invalid input")
		}

	}
	return "", bet
}
