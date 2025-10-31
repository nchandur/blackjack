package cmd

import (
	"fmt"
	"log"

	"github.com/nchandur/blackjack/player"
	"github.com/spf13/cobra"
)

var pm player.PlayerManager

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Start playing blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to Blackjack!!")
		return nil
	},
}

func Execute(playerManager player.PlayerManager) {
	pm = playerManager

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(playCmd)
	RootCmd.AddCommand(playerCmd)

	playerCmd.AddCommand(createCmd)
	playerCmd.AddCommand(signInCmd)
	playerCmd.AddCommand(retrieveCmd)
	playerCmd.AddCommand(deleteCmd)

}
