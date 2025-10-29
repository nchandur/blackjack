package models

import "fmt"

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

	for ok := true; ok; ok = d.FindSum() < 17 {

		fmt.Printf("%v\tSum: %d\n", d.Hand, d.Hand.FindSum())
		if d.CheckBlackjack() {
			fmt.Println("Blackjack!!!")
			break
		}

		if d.CheckBust() {
			fmt.Println("Bust :(")
			break
		}

		d.Hit(shoe)
	}

}
