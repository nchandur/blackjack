package player

import "fmt"

type Round struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
	Pushed int `json:"pushed"`
}

type Player struct {
	Username string `json:"username"`
	Password string
	SignedIn bool `json:"signedIn"`
	Round    `json:"rounds"`
}

func NewPlayer(username, password string) (Player, error) {
	if username == "" {
		return Player{}, fmt.Errorf("username cannot be empty")
	}

	if password == "" {
		return Player{}, fmt.Errorf("password cannot be empty")
	}

	return Player{
		Username: username,
		Password: password,
		SignedIn: false,
		Round:    Round{},
	}, nil

}

func (p *Player) String() string {

	return fmt.Sprintf("Username: %s\nSign In Status: %t\nStats: \nRounds Played: %d\nRounds Won: %d\nRounds Lost: %d\nRounds Pushed: %d", p.Username, p.SignedIn, p.Played, p.Won, p.Lost, p.Pushed)

}

type PlayerManager interface {
	SignIn(username, password string) error
	SignOut(username, password string) error
	Create(username, password string) error
	Retrieve(username string) (Player, error)
	Delete(username, password string) error
}
