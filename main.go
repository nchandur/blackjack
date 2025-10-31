package main

import (
	"log"

	"github.com/nchandur/blackjack/storage"
)

func main() {

	store, err := storage.Open("blackjack.json")

	if err != nil {
		log.Fatal(err)
	}

}
