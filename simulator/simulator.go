package simulator

import (
	"fmt"

	"github.com/nchandur/blackjack/entities/bot"
	"github.com/nchandur/blackjack/entities/dealer"
	"github.com/nchandur/blackjack/model"
)

type Simulator struct {
	bot.Bot
	dealer.Dealer
	shoe    *model.Shoe
	initial int64
}

// path to strategy.bin
func NewSimulator(strategy string) (*Simulator, error) {
	shoe := model.NewShoe()

	bot := bot.NewBot(&shoe)

	initial := bot.Kaasu

	err := bot.LoadStrats(strategy)

	if err != nil {
		return nil, fmt.Errorf("failed to create simulator: %v", err)
	}

	dealer := dealer.NewDealer(&shoe)

	return &Simulator{
		Bot:     *bot,
		Dealer:  *dealer,
		shoe:    &shoe,
		initial: initial,
	}, nil

}
