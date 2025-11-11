package main

import (
	"log"

	"github.com/nchandur/blackjack/utils"
)

func main() {

	out, err := utils.CsvToBin("./files/strategy/basic.csv")

	if err != nil {
		log.Fatal(err)
	}

	if err := utils.SaveBin(out, "./files/strategy/basic.bin"); err != nil {
		log.Fatal(err)
	}

}
