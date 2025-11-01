package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Hand
}

func (p *Player) Hit(shoe *Shoe) {
	p.Hand = append(p.Hand, shoe.Draw())
	p.Hand.MakeOptimum()
}

func (p *Player) CheckBust() bool {
	return p.Hand.FindSum() > 21
}

func (p *Player) CheckBlackjack() bool {
	return p.Hand.FindSum() == 21
}

func (p *Player) Play(shoe *Shoe, bet, funds int) (string, int) {

	reader := bufio.NewReader(os.Stdin)

	stop := false

	for ok := true; ok; ok = !stop {

		fmt.Printf("Player Hand\n%sSum: %d\n", p.Hand.String(), p.Hand.FindSum())

		if len(p.Hand) == 2 && p.CheckBlackjack() {
			bet = (2 * bet) + bet
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
