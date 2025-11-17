package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCSV(path string) ([][][]uint8, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to load strategy: %v", err)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to load strategy: %v", err)
	}

	ruleset := make([][][]uint8, 18)

	for i := range ruleset {
		ruleset[i] = make([][]uint8, 10)

		for j := range ruleset[i] {
			ruleset[i][j] = make([]uint8, 2)
		}

	}

	for _, record := range records[1:] {
		if len(record) != 4 {
			return nil, fmt.Errorf("failed to load strategy: malformed row")
		}

		hand, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, fmt.Errorf("failed to load strategy: %v", err)
		}

		dealer, err := strconv.Atoi(record[1])

		if err != nil {
			return nil, fmt.Errorf("failed to load strategy: %v", err)
		}

		isSoft, err := strconv.ParseBool(record[2])

		if err != nil {
			return nil, fmt.Errorf("failed to load strategy: %v", err)
		}

		action := strings.ToLower(record[3])[:1]

		var actionInt uint8

		switch action {
		case "h":
			actionInt = 1
		case "d":
			actionInt = 2
		default:
			actionInt = 0
		}

		if isSoft {
			ruleset[hand-4][dealer-2][1] = actionInt
		} else {
			ruleset[hand-4][dealer-2][0] = actionInt
		}

	}

	return ruleset, nil

}

func writeBIN(path string, data [][][]uint8) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := range data {
		for j := range data[i] {
			_, err = f.Write(data[i][j])

			if err != nil {
				return fmt.Errorf("failed to write strategy file: %v", err)
			}

		}
	}

	return nil
}

func main() {

	strats := []string{"basic", "h17", "s17", "enhc"}

	for _, strat := range strats {
		ruleset, err := readCSV(fmt.Sprintf("/home/chandur/workspace/goprojects/blackjack/files/strategy/%s.csv", strat))

		if err != nil {
			log.Fatal(err)
		}

		err = writeBIN(fmt.Sprintf("/home/chandur/workspace/goprojects/blackjack/files/strategy/%s.bin", strat), ruleset)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s.bin saved\n", strat)

	}

}
