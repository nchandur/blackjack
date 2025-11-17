package bot

import (
	"fmt"
	"os"

	"github.com/nchandur/blackjack/model"
)

type Bot struct {
	model.Hand
	ruleset [][][]uint8
	Kaasu   int64
	model.Stats
}

func NewBot(shoe *model.Shoe) *Bot {

	ruleset := make([][][]uint8, 18)

	for i := range ruleset {
		ruleset[i] = make([][]uint8, 10)

		for j := range ruleset[i] {
			ruleset[i][j] = make([]uint8, 2)
		}

	}

	return &Bot{Hand: model.NewHand(shoe), ruleset: ruleset, Kaasu: 1000}
}

func (b *Bot) LoadStrats(path string) error {

	f, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("failed to load strategy file: %v", err)
	}

	out := make([][][]uint8, 18)
	idx := 0
	for i := range 18 {
		out[i] = make([][]uint8, 10)
		for j := range 10 {
			out[i][j] = make([]uint8, 2)
			copy(out[i][j], f[idx:idx+2])
			idx += 2
		}
	}

	b.ruleset = out

	return nil
}

func (b *Bot) Lookup(dealer uint8) uint8 {

	if b.IsBlackjack() {
		return 0
	}

	if b.IsBust() {
		return 0
	}

	h := b.Sum()

	if b.IsSoft() {
		return b.ruleset[h-4][dealer-2][1]
	}

	return b.ruleset[h-4][dealer-2][0]

}
