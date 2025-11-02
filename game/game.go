package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nchandur/blackjack/game/dealer"
	"github.com/nchandur/blackjack/game/player"
	"github.com/nchandur/blackjack/game/structs"
)

const SHOE_SIZE int = 6

type Game struct {
	player.Player
	dealer.Dealer
	structs.Shoe
	Rounds int
	Wins   int
	Losses int
	Pushes int
	Kaasu  int
}

func NewGame() Game {
	shoe := structs.NewShoe(SHOE_SIZE)

	player := player.Player{Hand: structs.NewHand(&shoe)}
	dealer := dealer.Dealer{Hand: structs.NewHand(&shoe)}

	return Game{
		Player: player,
		Dealer: dealer,
		Shoe:   shoe,
		Rounds: 0,
		Wins:   0,
		Losses: 0,
		Pushes: 0,
		Kaasu:  0,
	}

}

func (g *Game) Play() error {

	for {
		action, err := g.playRound()

		if err != nil {
			log.Println(err)
		}

		if action == "q" {
			break
		}

	}

	return nil

}

func (g *Game) playRound() (string, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %d kaasu.\nEnter bet amount: ", g.Kaasu)
	input, _ := reader.ReadString('\n')

	input = strings.Trim(input, "\n")

	if input == "q" {
		return "q", nil
	}

	bet, err := strconv.Atoi(input)

	if err != nil {
		return "q", fmt.Errorf("failed to play game: invalid bet amount: %v", err)
	}

	if bet > g.Kaasu {
		return "q", fmt.Errorf("failed to play game: insufficient funds")
	}

	fmt.Printf("Dealer Hand\n%s\n", strings.Join(g.Dealer.Hand[0].String(), "\n"))

	action, bet := g.Player.Play(reader, &g.Shoe, bet, g.Kaasu)

	if action == "q" {
		return action, nil
	}

	g.Dealer.Play(&g.Shoe)

	fmt.Printf("Final\nDealer Hand\n%sPlayer Hand\n%s", g.Dealer.Hand.String(), g.Player.Hand.String())

	switch g.outcome() {
	case 1:
		fmt.Printf("Player Won :)\n\n")
		g.Wins++
		g.Kaasu += bet
	case -1:
		fmt.Printf("Player Lost :(\n\n")
		g.Kaasu -= bet
		g.Losses++
	case 0:
		fmt.Printf("Push\n\n")
		g.Pushes++
	}

	g.Rounds++

	g.Player.Hand = structs.NewHand(&g.Shoe)
	g.Dealer.Hand = structs.NewHand(&g.Shoe)

	g.reset()

	return action, nil

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

func (g *Game) reset() {
	left := 0
	total := SHOE_SIZE * 52

	for _, val := range g.Shoe {
		left += val
	}

	if float64(left)/float64(total) < 0.2 {
		g.Shoe = structs.NewShoe(SHOE_SIZE)
	}

}
