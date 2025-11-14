package model

import "fmt"

type Card struct {
	value uint8
	rank  string
	suit  string
}

func (c *Card) String() []string {
	return []string{
		"-----",
		fmt.Sprintf("│%-2s%s│", c.rank, c.suit),
		"-----",
	}
}
