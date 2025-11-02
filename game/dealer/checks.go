package dealer

func (d *Dealer) CheckBust() bool {
	return d.Hand.FindSum() > 21
}

func (d *Dealer) CheckBlackjack() bool {
	return d.Hand.FindSum() == 21
}
