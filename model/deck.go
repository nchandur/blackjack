package model

func newDeck() []Card {

	deck := []Card{}

	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	values := []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

	for i, rank := range ranks {
		for _, suit := range []string{"C", "H", "S", "D"} {
			card := Card{Value: values[i], rank: rank, suit: suit}
			deck = append(deck, card)
		}

	}

	return deck

}
