package main

import (
	"fmt"

	"github.com/nchandur/blackjack/models"
)

func main() {

	shoe := models.NewShoe(5)

	player := models.Player{Hand: models.NewHand(&shoe)}
	dealer := models.Dealer{Hand: models.NewHand(&shoe)}

	fmt.Printf("Dealer Hand: %s\n", dealer.Hand[0].String())

	player.Play(&shoe)
	dealer.Play(&shoe)

	fmt.Println("Player Hand: ", player.String(), player.FindSum())
	fmt.Println("Dealer Hand: ", dealer.String(), dealer.FindSum())

}
