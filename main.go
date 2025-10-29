package main

import (
	"fmt"

	"github.com/nchandur/blackjack/models"
)

func main() {

	shoe := models.NewShoe(5)

	player := models.Player{Hand: models.NewHand(&shoe)}
	dealer := models.Dealer{Hand: models.NewHand(&shoe)}

	fmt.Println("Player Hand: ", player.Hand, player.Hand.FindSum())

	fmt.Println("Dealer Hand: ", dealer.Hand[0])

	player.Play(&shoe)

	dealer.Play(&shoe)

}
