package storage

import (
	"fmt"

	"github.com/nchandur/blackjack/player"
)

type PlayerManager struct {
	players *[]player.Player
}

func NewPlayerManager(players *[]player.Player) PlayerManager {
	return PlayerManager{players: players}
}

func (p PlayerManager) Retrieve(username string) (player.Player, error) {
	for _, play := range *(p.players) {
		if play.Username == username {
			return play, nil
		}
	}

	return player.Player{}, fmt.Errorf("failed to retrieve player: player not found")

}

func (p PlayerManager) Create(username, password string) error {
	_, err := p.Retrieve(username)

	if err == nil {
		return fmt.Errorf("failed to create player: username %s already exists", username)
	}

	play, err := player.NewPlayer(username, password)

	if err != nil {
		return fmt.Errorf("failed to create player: %v", err)
	}

	*(p.players) = append(*(p.players), play)

	return nil

}

func (p PlayerManager) Delete(username, password string) error {
	idx := -1

	for i, play := range *(p.players) {
		if play.Username == username {
			if play.Password == password {
				idx = i
				break
			} else {
				return fmt.Errorf("failed to delete player: invalid password")
			}
		}
	}

	if idx == -1 {
		return fmt.Errorf("failed to delete player: %s does not exist", username)
	}

	*(p.players) = append((*(p.players))[:idx], (*(p.players))[idx+1:]...)

	return nil
}

func (p PlayerManager) SignIn(username, password string) error {

	for idx := range *(p.players) {
		if (*(p.players))[idx].Username == username && (*(p.players))[idx].Password == password {
			(*(p.players))[idx].SignedIn = true
			return nil
		} else {
			(*(p.players))[idx].SignedIn = false

		}
	}

	return fmt.Errorf("username or password is incorrect")

}

func (p PlayerManager) SignOut(username, password string) error {

	for idx := range *(p.players) {
		if (*(p.players))[idx].Username == username && (*(p.players))[idx].Password == password {
			(*(p.players))[idx].SignedIn = false
			return nil
		}
	}

	return fmt.Errorf("username or password is incorrect")

}
