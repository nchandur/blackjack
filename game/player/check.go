package player

func (p *Player) CheckBust() bool {
	return p.Hand.FindSum() > 21
}

func (p *Player) CheckBlackjack() bool {
	return p.Hand.FindSum() == 21
}
