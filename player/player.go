package player

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const BUY_IN int = 1000

type Round struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
	Pushed int `json:"pushed"`
}

type Player struct {
	Username  string `json:"username"`
	Password  string
	SignedIn  bool `json:"signedIn"`
	Round     `json:"rounds"`
	Buckaroos int `json:"buckaroos"`
}

func NewPlayer(username, password string) (Player, error) {
	if username == "" {
		return Player{}, fmt.Errorf("failed to create player: username cannot be empty")
	}

	if password == "" {
		return Player{}, fmt.Errorf("failed to create player: password cannot be empty")
	}

	play := Player{
		Username:  username,
		SignedIn:  false,
		Round:     Round{},
		Buckaroos: BUY_IN,
	}

	hash, err := play.encryptPassword(password)

	if err != nil {
		return Player{}, fmt.Errorf("failed to create player: %v", err)
	}

	play.Password = hash

	return play, nil

}

func (p Player) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}

func (p Player) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	return err == nil
}

func (p *Player) String() string {

	return fmt.Sprintf("Username: %s\nSign In Status: %t\n\nStats: \nRounds Played: %d\nRounds Won: %d\nRounds Lost: %d\nRounds Pushed: %d\n\nBuckaroos: %d", p.Username, p.SignedIn, p.Played, p.Won, p.Lost, p.Pushed, p.Buckaroos)

}

type PlayerManager interface {
	SignIn(username, password string) error
	SignOut(username, password string) error
	Create(username, password string) error
	Retrieve(username string) (Player, error)
	Delete(username, password string) error
}
