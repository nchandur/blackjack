package models

type Hand []Card

func NewHand(shoe *Shoe) Hand {
	hand := make(Hand, 0)

	hand = append(hand, shoe.Draw())
	hand = append(hand, shoe.Draw())

	hand.makeOptimum()

	return hand
}

func (h *Hand) CheckBust() bool {

	return h.FindSum() > 21
}

func (h *Hand) CheckBlackjack() bool {
	return h.FindSum() == 21

}

func (h *Hand) FindSum() int {
	sum := 0

	for _, card := range *h {
		sum += card.Value
	}

	return sum

}

func (h *Hand) makeOptimum() {

	for idx := range *h {
		if h.FindSum() > 21 && (*h)[idx].Rank == "A" {
			(*h)[idx].Value = 1
		}
	}

}

func (h *Hand) Hit(shoe *Shoe) {
	(*h) = append((*h), (*shoe).Draw())
	h.makeOptimum()
}
