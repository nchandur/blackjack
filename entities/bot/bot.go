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

func (b *Bot) Hit(shoe *structs.Shoe) {
	(*b).Hand = append((*b).Hand, shoe.Draw())
}

func (b *Bot) CheckBlackjack() bool {
	return b.Hand.FindSum() == 21
}

func (b *Bot) CheckBust() bool {
	return b.Hand.FindSum() > 21
}
