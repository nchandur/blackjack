package models

type Deck []Card

func NewDeck() Deck {

	deck := []Card{}

	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

	for i, rank := range ranks {
		for _, suit := range []string{"c", "h", "s", "d"} {
			card := Card{Value: values[i], Rank: rank, Suit: suit}
			deck = append(deck, card)
		}

	}

	return deck

}
