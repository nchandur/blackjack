package cmd

import (
	"github.com/spf13/cobra"
)

var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Play Blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
