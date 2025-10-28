package models

type Hand []Card

func NewHand(shoe *Shoe) Hand {
	hand := make(Hand, 0)

	hand = append(hand, shoe.Draw())
	hand = append(hand, shoe.Draw())

	return hand
}

func (h *Hand) CheckBust() bool {

	return h.findSum() > 21
}

func (h *Hand) CheckBlackjack() bool {
	return h.findSum() == 21

}

func (h *Hand) findSum() int {
	sum := 0

	for _, card := range *h {
		sum += card.Value
	}

	return sum

}

func (h *Hand) MakeOptimum() {

}
