package utils

import (
	"encoding/binary"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/nchandur/blackjack/entities/bot"
)

func CsvToBin(path string) ([]uint16, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	encoded := make([]uint16, 0)

	for _, record := range records[1:] {

		botHand, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, err
		}

		dealerCard, err := strconv.Atoi(record[1])

		if err != nil {
			return nil, err
		}

		isSoft, err := strconv.ParseBool(record[2])

		if err != nil {
			return nil, err
		}

		action := byte(record[3][0])

		rule := bot.Rule{
			BotHand:    uint16(botHand),
			DealerCard: uint16(dealerCard),
			IsSoft:     isSoft,
			Action:     action,
		}

		bits := BitEncoder(rule)

		encoded = append(encoded, bits)

	}

	return encoded, nil
}

func SaveBin(bits []uint16, destination string) error {
	f, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, bits)
	return err

}

func boolToInt(flag bool) uint16 {
	if flag {
		return 1
	}
	return 0
}

func intToBool(val uint16) bool {
	return val != 0
}

func BitEncoder(rule bot.Rule) uint16 {

	botHand := (rule.BotHand - 4) << 11
	dealerCard := (rule.DealerCard - 2) << 7
	isSoft := (boolToInt(rule.IsSoft)) << 6
	action := bot.ACTION_MAP[rule.Action] << 4

	return uint16(botHand) | uint16(dealerCard) | isSoft | action
}

func RuleDecoder(bits uint16) bot.Rule {

	rule := bot.Rule{}

	rule.BotHand = (bits >> 11 & 0b11111) + 4
	rule.DealerCard = (bits >> 7 & 0b1111) + 2
	rule.IsSoft = intToBool(bits >> 6 & 0b1)
	rule.Action = bot.REVERSE_ACTION_MAP[bits>>4&0b11]

	return rule
}
