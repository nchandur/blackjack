package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nchandur/blackjack/users"
)

type Storage struct {
	Filename string
	gamers   []users.User
}

func Open(filename string) (Storage, error) {
	file, err := os.Open(filename)

	if err != nil {
		return Storage{gamers: nil}, fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return Storage{gamers: nil}, fmt.Errorf("error retrieving data: %v", err)
	}

	var storage Storage

	err = json.Unmarshal(bytes, &storage.gamers)

	if err != nil {
		return Storage{gamers: nil}, fmt.Errorf("error unmarshalling json: %v", err)
	}

	storage.Filename = filename
	return storage, nil
}

func (s *Storage) Retrieve(username string) 