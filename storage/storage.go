package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nchandur/blackjack/player"
)

type Storage struct {
	Filename string
}

func NewStorage(filename string) *Storage {
	storage := Storage{Filename: filename}

	return &storage
}

func (s *Storage) Save(players []player.Player) error {

	fileData, err := json.MarshalIndent(players, "", "\t")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, fileData, 0644)

}

func (s *Storage) Load(players *[]player.Player) error {

	fileData, err := os.ReadFile(s.Filename)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("failed to read file '%s': %w", s.Filename, err)
		}
		return fmt.Errorf("failed to read file '%s': %w", s.Filename, err)
	}

	if len(fileData) == 0 {
		return fmt.Errorf("failed to read file '%s': no bytes received", s.Filename)
	}

	err = json.Unmarshal(fileData, &players)

	if err != nil {
		return fmt.Errorf("failed to unmarshal data from '%s': %w", s.Filename, err)
	}

	return nil
}
