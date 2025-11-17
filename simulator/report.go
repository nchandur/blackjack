package simulator

import "fmt"

func (s *Simulator) Report() string {
	return fmt.Sprintf("%s\nInitial: %d\tFinal: %d\nNet: %d\n", s.Bot.Stats.String(), s.initial, s.Bot.Kaasu, s.Bot.Kaasu-s.initial)
}
