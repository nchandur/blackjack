package main

import (
	"fmt"

	"github.com/nchandur/blackjack/models"
)

func main() {

	shoe := models.NewShoe(10)

	playerHand := models.NewHand(&shoe)

	fmt.Println(playerHand)

}
