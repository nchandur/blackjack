package game

import (
	"fmt"
	"strings"
)

const SHOE_SIZE int = 6

type Game struct {
	Player
	Dealer
	Shoe
	Rounds int
	Wins   int
	Losses int
	Pushes int
}

func NewGame() Game {
	shoe := NewShoe(SHOE_SIZE)

	player := Player{NewHand(&shoe)}
	dealer := Dealer{NewHand(&shoe)}

	return Game{
		Player: player,
		Dealer: dealer,
		Shoe:   shoe,
		Rounds: 0,
		Wins:   0,
		Losses: 0,
		Pushes: 0,
	}

}

func (g *Game) Play() Game {

	for {
		action := g.playRound()

		if action == "q" {
			break
		}

	}

	return *g

}

func (g *Game) playRound() string {

	fmt.Printf("Dealer Hand\n%s\n", strings.Join(g.Dealer.Hand[0].String(), "\n"))

	action := g.Player.Play(&g.Shoe)

	if action == "q" {
		return action
	}

	g.Dealer.Play(&g.Shoe)

	fmt.Printf("Final\nDealer Hand\n%sPlayer Hand\n%s", g.Dealer.Hand.String(), g.Player.Hand.String())

	switch g.outcome() {
	case 1:
		fmt.Printf("Player Won :)\n\n")
		g.Wins++
	case -1:
		fmt.Printf("Player Lost :(\n\n")
		g.Losses++
	case 0:
		fmt.Printf("Push\n\n")
		g.Pushes++
	}

	g.Rounds++

	g.Player.Hand = NewHand(&g.Shoe)
	g.Dealer.Hand = NewHand(&g.Shoe)

	g.reset()

	return action

}

func (g *Game) outcome() int32 {

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

func (g *Game) GetReport() string {

	return fmt.Sprintf("Rounds played: %d\nRounds won: %d\nRounds Lost: %d\nRounds Pushed: %d\n", g.Rounds, g.Wins, g.Losses, g.Pushes)

}

func (g *Game) reset() {
	left := 0
	total := SHOE_SIZE * 52

	for _, val := range g.Shoe {
		left += val
	}

	if float64(left)/float64(total) < 0.2 {
		g.Shoe = NewShoe(SHOE_SIZE)
	}

}
