package model

import (
	"strings"
)

type Hand []Card

func NewHand(shoe *Shoe) Hand {
	hand := make(Hand, 0)

	hand = append(hand, shoe.Draw())
	hand = append(hand, shoe.Draw())

	return hand
}

func (h *Hand) isSoft() bool {
	for _, card := range *h {
		if card.rank == "A" {
			return true
		}
	}
	return false
}

func (h *Hand) Sum() uint8 {
	var sum uint8

	for _, card := range *h {
		sum += card.value
	}

	if sum > 21 && h.isSoft() {
		sum -= 10
	}

	return sum

}

func (h *Hand) Hit(shoe *Shoe) {
	(*h) = append((*h), shoe.Draw())
}

func (h *Hand) IsBlackjack() bool {
	return h.Sum() == 21
}

func (h *Hand) IsBust() bool {
	return h.Sum() > 21
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
