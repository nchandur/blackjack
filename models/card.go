package models

import "fmt"

type Card struct {
	Value int
	Rank  string
	Suit  string
}

func (c *Card) String() []string {
	return []string{
		"┌───┐",
		fmt.Sprintf("│%-2s%s│", c.Rank, c.Suit),
		"└───┘",
	}
}
