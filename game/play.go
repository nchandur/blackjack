package game

import (
	"fmt"
	"strings"

	"github.com/nchandur/blackjack/game/structs"
)

func (g *Game) Play() {
	before := 0

	for {

		fmt.Printf("----------------------------------New Round----------------------------------\n")
		fmt.Printf("You have %d kaasu.\nEnter bet amount: ", g.Player.Kaasu)
		_, err := fmt.Scanln(&before)

		if err != nil {
			break
		}

		if before > g.Player.Kaasu {
			break
		}

		fmt.Printf("Dealer Hand: \n%s", strings.Join(g.Dealer.Hand[0].String(), "\n"))

		after := g.Player.Play(g.shoe, before)

		if after == -1 {
			break
		}

		after = g.Dealer.Play(g.shoe, after)

		out := g.outcome()

		switch out {
		case "WON":
			g.Player.Won++
			g.Dealer.Lost++
			g.Player.Kaasu += after
			g.Dealer.Kaasu -= after
		case "LOST":
			g.Player.Lost++
			g.Dealer.Won++
			g.Player.Kaasu -= after
			g.Dealer.Kaasu += after
		}

		fmt.Printf("-----------------------------------Outcome-----------------------------------\nDealer Hand: %d\n%s\nPlayer Hand: %d\n%s\n%s\n", g.Dealer.FindSum(), &g.Dealer.Hand, g.Player.FindSum(), &g.Player.Hand, out)

		g.Player.Played++
		g.Dealer.Played++

		g.Player.Hand = structs.NewHand(g.shoe)
		g.Dealer.Hand = structs.NewHand(g.shoe)

		g.reset()

	}

}
