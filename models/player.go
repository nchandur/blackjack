package models

import (
	"bufio"
	"fmt"
	"log"
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

func (p *Player) Play(shoe *Shoe) {

	reader := bufio.NewReader(os.Stdin)

	stop := false

	for ok := true; ok; ok = !stop {

		fmt.Printf("Player Hand:\n%s\tSum: %d\n", p.String(), p.FindSum())

		if p.CheckBlackjack() {
			break
		}

		if p.CheckBust() {
			break
		}

		fmt.Printf("Your move (h/s/d): ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
		}

		input = strings.ToLower(strings.Trim(input, "\n"))

		switch input {
		case "h", "hit":
			p.Hit(shoe)
		case "s", "stand":
			stop = true
		case "d", "double":
			p.Hit(shoe)
			stop = true
		default:
			fmt.Println("invalid input")
		}

	}
}
