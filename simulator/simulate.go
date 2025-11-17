package simulator

import (
	"github.com/nchandur/blackjack/model"
)

func (s *Simulator) isWinner() int8 {
	if s.Bot.Hand.IsBust() {
		return -1
	}

	if s.Bot.Hand.IsBlackjack() {
		return 1
	}

	if s.Dealer.Hand.IsBust() {
		return 1
	}

	if s.Dealer.Hand.IsBlackjack() {
		return -1
	}

	if s.Dealer.Hand.Sum() == s.Bot.Sum() {
		return 0
	}

	if s.Bot.Sum() > s.Dealer.Sum() {
		return 1
	}

	return -1

}

func (s *Simulator) Simulate(rounds uint64) {
	initial := 1
	bet := int64(initial)

	count := uint64(1)

	for count <= rounds {

		after := s.Bot.Play(bet, s.Dealer.Hand[0].Value, s.shoe)
		after = s.Dealer.Play(after, s.shoe)

		switch s.isWinner() {
		case 0:
			s.Bot.Stats.Pushed++
			s.Dealer.Stats.Pushed++
		case 1:
			s.Bot.Kaasu += after
			s.Dealer.Kaasu -= after
			s.Bot.Stats.Won++
			s.Dealer.Stats.Lost++
			bet = (bet) * (int64(count) % 4)
		case -1:
			s.Bot.Kaasu -= after
			s.Dealer.Kaasu += after
			s.Bot.Stats.Lost++
			s.Dealer.Stats.Won++
			bet = int64(initial)
		}

		s.Bot.Stats.Played++
		s.Dealer.Stats.Played++

		s.Bot.Hand = model.NewHand(s.shoe)
		s.Dealer.Hand = model.NewHand(s.shoe)

		s.shoe.Reset()

		count++
	}
}
