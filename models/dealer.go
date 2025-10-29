package models

type Dealer struct {
	Hand
}

func (d *Dealer) Hit(shoe *Shoe) {
	d.Hand = append(d.Hand, shoe.Draw())
	d.Hand.MakeOptimum()
}

func (d *Dealer) CheckBust() bool {
	return d.Hand.FindSum() > 21
}

func (d *Dealer) CheckBlackjack() bool {
	return d.Hand.FindSum() == 21
}

func (d *Dealer) Play(shoe *Shoe) {

	for d.FindSum() <= 17 {

		if d.CheckBlackjack() {
			break
		}

		if d.CheckBust() {
			break
		}

		d.Hit(shoe)
	}

}
