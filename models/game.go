package models

import "fmt"

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

func (g *Game) Play() {

	for {
		action := g.PlayRound()

		if action == "q" {
			break
		}

	}

}

func (g *Game) PlayRound() string {

	fmt.Printf("Dealer Card:\n%s", g.Dealer.Hand.String())

	action := g.Player.Play(&g.Shoe)

	if action == "q" {
		return action
	}

	g.Dealer.Play(&g.Shoe)

	fmt.Printf("\n--------------------------------------------------------------------------\n")
	fmt.Printf("Dealer Hand: \n%s", g.Dealer.Hand.String())
	fmt.Printf("Player Hand: \n%s", g.Player.Hand.String())

	switch g.didPlayerWin() {
	case 1:
		fmt.Println("Player Won :)")
		g.Wins++
	case -1:
		fmt.Println("Player Lost :(")
	case 0:
		fmt.Println("Push")
	}

	g.Rounds++

	g.Player.Hand = NewHand(&g.Shoe)
	g.Dealer.Hand = NewHand(&g.Shoe)

	return action

}

func (g *Game) didPlayerWin() int32 {

	if g.Player.CheckBust() {
		return -1
	}

	if g.Dealer.CheckBust() {
		return 1
	}

	if g.Player.FindSum() == g.Dealer.FindSum() {
		return 0
	}

	if g.Player.CheckBlackjack() {
		return 1
	}

	if g.Dealer.CheckBlackjack() {
		return -1
	}

	if g.Player.FindSum() > g.Dealer.FindSum() {
		return 1
	}

	return -1

}
