package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Play Blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play Blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		g, err := gm.Load()

		if err != nil {
			return err
		}

		g.Play()

		err = gm.Save(g)

		if err != nil {
			return fmt.Errorf("failed to save game: %v", err)
		}

		return nil

	},
}
