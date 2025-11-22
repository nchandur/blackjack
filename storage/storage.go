package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nchandur/blackjack/model"
	"github.com/nchandur/blackjack/users"
)

type Storage struct {
	Filename string
}

func NewStorage(filename string) *Storage {
	storage := Storage{Filename: filename}

	return &storage
}

func (s *Storage) Save(players []users.User) error {

	fileData, err := json.MarshalIndent(players, "", "\t")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, fileData, 0644)

}

func (s *Storage) Load(players *[]users.User) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(s.Filename)
			if createErr != nil {
				return fmt.Errorf("failed to create file '%s': %w", s.Filename, createErr)
			}

			defer file.Close()

			*players = []users.User{{Username: "HOUSE", Password: "", SignedIn: false, Stats: model.Stats{}, Kaasu: 0}}
			return os.WriteFile(s.Filename, fileData, 0644)
		}
		return fmt.Errorf("failed to read file '%s': %w", s.Filename, err)
	}

	if len(fileData) == 0 {
		*players = []users.User{{Username: "HOUSE", Password: "", SignedIn: false, Stats: model.Stats{}, Kaasu: 0}}
		return os.WriteFile(s.Filename, fileData, 0644)
	}

	if err := json.Unmarshal(fileData, players); err != nil {
		return fmt.Errorf("failed to unmarshal data from '%s': %w", s.Filename, err)
	}

	return nil
}
