package models

import "fmt"

type Card struct {
	Value int
	Rank  string
	Suit  string
}

func (c *Card) String() string {
	return fmt.Sprintf("[%s%s]", c.Rank, c.Suit)
}
