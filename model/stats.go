package model

import "fmt"

type Stats struct {
	Played uint64
	Won    uint64
	Lost   uint64
	Pushed uint64
}

func (s *Stats) String() string {
	return fmt.Sprintf("Stats\nPlayed: %d\nWon: %d\nLost: %d\nPushed: %d\n", s.Played, s.Won, s.Lost, s.Pushed)
}
