package models

import (
	"fmt"
)

type Game struct {
	Player
	Dealer
	Shoe
	Rounds int
	Wins   int
}

func NewGame() Game {
	shoe := NewShoe(6)

	player := Player{NewHand(&shoe)}
	dealer := Dealer{NewHand(&shoe)}

	return Game{
		Player: player,
		Dealer: dealer,
		Shoe:   shoe,
		Rounds: 0,
		Wins:   0,
	}

}

func (g *Game) PlayRound() {

	fmt.Printf("Dealer Hand: \n%s\n", g.Dealer.Hand.String())

	g.Player.Play(&g.Shoe)
	g.Dealer.Play(&g.Shoe)

	fmt.Printf("Player Hand:\n%s\nSum: %d\n", g.Player.String(), g.Player.FindSum())
	fmt.Printf("Dealer Hand:\n%s\nSum: %d\n", g.Dealer.String(), g.Dealer.FindSum())

	if g.didPlayerWin() {
		fmt.Println("You won!!")
		return
	}

	fmt.Println("You lost :(")

}

func (g *Game) didPlayerWin() bool {

	if g.Player.CheckBlackjack() {
		return true
	}

	if g.Player.CheckBust() {
		return false
	}

	if g.Dealer.CheckBlackjack() {
		return false
	}

	if g.Dealer.CheckBust() {
		return true
	}

	return g.Player.FindSum() > g.Dealer.FindSum()

}
