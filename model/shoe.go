package model

import (
	"fmt"
	"math/rand/v2"
)

const SHOE_SIZE = 6

type Shoe map[Card]int

func NewShoe() Shoe {

	shoe := make(map[Card]int)
	deck := newDeck()

	for _, card := range deck {
		shoe[card] = SHOE_SIZE
	}

	return shoe
}

func (s *Shoe) Draw() Card {

	left := make([]Card, 0)

	for key, val := range *s {
		if val != 0 {
			left = append(left, key)
		}
	}

	card := left[rand.IntN(len(left))]
	(*s)[card]--

	return card
}

func (s *Shoe) cardCount() int {
	sum := 0

	for _, val := range *s {
		sum += val
	}

	return sum

}

func (s *Shoe) Reset() {

	left := s.cardCount()

	if float32(left)/float32((SHOE_SIZE*52)) < float32(0.2) {
		*s = NewShoe()
	}

}

func (s *Shoe) Display() {
	for card, left := range *s {
		fmt.Printf("Card: %v\tLeft: %d\n", card, left)
	}

}
