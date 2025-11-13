package structs

import (
	"strings"
)

type Hand []Card

func NewHand(shoe *Shoe) Hand {
	hand := make(Hand, 0)

	hand = append(hand, shoe.Draw())
	hand = append(hand, shoe.Draw())

	hand.MakeOptimum()

	return hand
}

func (h *Hand) FindSum() int {
	sum := 0

	for _, card := range *h {
		sum += card.Value
	}

	return sum

}

func (h *Hand) MakeOptimum() {

	for idx := range *h {
		if h.FindSum() > 21 && (*h)[idx].Rank == "A" {
			(*h)[idx].Value = 1
		}
	}

}

func (h *Hand) String() string {

	var builder strings.Builder
	var cardLines [][]string

	for _, c := range *h {
		cardLines = append(cardLines, c.String())
	}

	for i := range cardLines[0] {
		for _, card := range cardLines {
			builder.WriteString(card[i])
			builder.WriteString("  ")
		}
		builder.WriteString("\n")
	}

	return builder.String()

}
