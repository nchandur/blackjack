package bot

import (
	"github.com/nchandur/blackjack/entities"
	"github.com/nchandur/blackjack/game/structs"
)

type Bot struct {
	structs.Hand
	Kaasu int
	entities.Stats
}

func NewBot(shoe *structs.Shoe) *Bot {
	return &Bot{
		Hand: structs.NewHand(shoe),
	}
}
