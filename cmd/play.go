package cmd

import (
	"github.com/nchandur/blackjack/game"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play Blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := game.NewGame()

		g.Play()

		return nil

	},
}
