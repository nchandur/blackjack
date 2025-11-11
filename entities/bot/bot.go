package bot

import (
	"github.com/nchandur/blackjack/entities"
	"github.com/nchandur/blackjack/game/structs"
)

var ACTION_MAP = map[byte]uint16{
	's': 0,
	'h': 1,
	'd': 2,
}

var REVERSE_ACTION_MAP = map[uint16]byte{
	0: 's',
	1: 'h',
	2: 'd',
}

type Rule struct {
	BotHand    uint16
	DealerCard uint16
	IsSoft     bool
	Action     byte
}

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
