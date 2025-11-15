package main

import (
	"fmt"
	"log"

	"github.com/nchandur/blackjack/bot"
	"github.com/nchandur/blackjack/dealer"
	"github.com/nchandur/blackjack/model"
)

func isWinner(bot *bot.Bot, dealer *dealer.Dealer) int8 {
	if bot.Hand.IsBust() {
		return -1
	}

	if bot.Hand.IsBlackjack() {
		return 1
	}

	if dealer.Hand.IsBust() {
		return 1
	}

	if dealer.Hand.IsBlackjack() {
		return -1
	}

	if dealer.Hand.Sum() == bot.Sum() {
		return 0
	}

	if bot.Sum() > dealer.Sum() {
		return 1
	}

	return -1

}

func Simulate(bot *bot.Bot, dealer *dealer.Dealer, shoe *model.Shoe, rounds uint64) {
	bet := int64(100)

	count := uint64(0)

	for count < rounds {

		after := bot.Play(bet, dealer.Hand[0].Value, shoe)
		after = dealer.Play(after, shoe)

		switch isWinner(bot, dealer) {
		case 0:
			bot.Stats.Pushed++
			dealer.Stats.Pushed++
		case 1:
			bot.Kaasu += after
			dealer.Kaasu -= after
			bot.Stats.Won++
			dealer.Stats.Lost++
		case -1:
			bot.Kaasu -= after
			dealer.Kaasu += after
			bot.Stats.Lost++
			dealer.Stats.Won++
		}

		bot.Stats.Played++
		dealer.Stats.Played++

		bot.Hand = model.NewHand(shoe)
		dealer.Hand = model.NewHand(shoe)

		shoe.Reset()

		count++
	}
}

func main() {

	shoe := model.NewShoe()

	dealer := dealer.NewDealer(&shoe)
	bot := bot.NewBot(&shoe)

	err := bot.LoadStrats("./files/strategy/basic.bin")

	if err != nil {
		log.Fatal(err)
	}

	initial := bot.Kaasu

	fmt.Printf("Initial Funds: %d\n\n", initial)

	Simulate(bot, dealer, &shoe, 1000)

	fmt.Printf("%s\nFinal Funds: %d\nNet Funds: %d\n", bot.Stats.String(), bot.Kaasu, bot.Kaasu-initial)

}
