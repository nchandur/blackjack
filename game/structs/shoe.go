package structs

import (
	"fmt"
	"math/rand/v2"
)

type Shoe map[Card]int

func NewShoe(n int) Shoe {
	shoe := make(map[Card]int)

	deck := NewDeck()

	for _, card := range deck {
		shoe[card] = n
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

func (s *Shoe) Display() {
	for card, left := range *s {
		fmt.Printf("Card: %v\tLeft: %d\n", card, left)
	}

}

func (s *Shoe) GetRemainingCards() int {
	sum := 0

	for _, val := range *s {
		sum += val
	}

	return sum

}

func (s *Shoe) Reset(n int) Shoe {
	return NewShoe(n)
}
