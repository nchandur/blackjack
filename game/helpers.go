package game

import "github.com/nchandur/blackjack/game/structs"

func (g *Game) outcome() string {

	if g.Player.CheckBust() {
		return "LOST"
	}

	if g.Dealer.CheckBust() {
		return "WON"
	}

	if g.Player.FindSum() == g.Dealer.FindSum() {
		return "PUSH"
	}

	if g.Player.CheckBlackjack() {
		return "WON"
	}

	if g.Dealer.CheckBlackjack() {
		return "LOST"
	}

	if g.Player.FindSum() > g.Dealer.FindSum() {
		return "WON"
	}

	return "LOST"

}

func (g *Game) reset() {
	left := 0
	total := SHOE_SIZE * 52

	for _, val := range *g.shoe {
		left += val
	}

	if float64(left)/float64(total) < 0.2 {
		*g.shoe = structs.NewShoe(SHOE_SIZE)
	}

}
