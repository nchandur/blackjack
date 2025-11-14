package game

import (
	"fmt"
	"strings"

	"github.com/nchandur/blackjack/model"
)

func (g *Game) Play() error {

	for {

		var bet int64
		fmt.Printf("\n-----------------------------------------New Round-----------------------------------------\nYou have %d kaasu. Place your bet: ", g.Player.Kaasu)
		_, err := fmt.Scanln(&bet)

		if err != nil {
			return nil
		}

		if bet > g.Player.Kaasu {
			return fmt.Errorf("failed to play: insufficient funds")
		}

		if bet == 0 {
			return fmt.Errorf("failed to play: bet amount must be non-zero")
		}

		fmt.Printf("Dealer Hand: \n%s\n", strings.Join(g.Dealer.Hand[0].String(), "\n"))

		after := g.Player.Play(bet, &g.shoe)

		if after == 0 {
			return nil
		}

		after = g.Dealer.Play(after, &g.shoe)

		fmt.Printf("\n------------------------------------------Outcome------------------------------------------\nDealer Hand: %d\n%s\nPlayer Hand: %d\n%s\n", g.Dealer.Sum(), g.Dealer.Hand.String(), g.Player.Sum(), g.Player.Hand.String())

		switch g.isWinner() {
		case 0:
			fmt.Println("PUSH")
			g.Player.Stats.Pushed++
			g.Dealer.Stats.Pushed++
		case 1:
			fmt.Println("WON")
			g.Player.Kaasu += after
			g.Dealer.Kaasu -= after
			g.Player.Stats.Won++
			g.Dealer.Stats.Lost++
		case -1:
			fmt.Println("LOST")
			g.Player.Kaasu -= after
			g.Dealer.Kaasu += after
			g.Player.Stats.Lost++
			g.Dealer.Stats.Won++
		}

		g.Player.Stats.Played++
		g.Dealer.Stats.Played++

		g.Player.Hand = model.NewHand(&g.shoe)
		g.Dealer.Hand = model.NewHand(&g.shoe)

		g.shoe.Reset()

	}

}

func (g *Game) isWinner() int8 {
	if g.Player.Hand.IsBust() {
		return -1
	}

	if g.Player.Hand.IsBlackjack() {
		return 1
	}

	if g.Dealer.Hand.IsBust() {
		return 1
	}

	if g.Dealer.Hand.IsBlackjack() {
		return -1
	}

	if g.Dealer.Hand.Sum() == g.Player.Sum() {
		return 0
	}

	if g.Player.Sum() > g.Dealer.Sum() {
		return 1
	}

	return -1
}
