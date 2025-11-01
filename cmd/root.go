package cmd

import (
	"fmt"
	"log"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/player"
	"github.com/spf13/cobra"
)

var pm player.PlayerManager
var gm game.GameManager

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Start playing blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to Blackjack!!")
		return nil
	},
}

func Execute(playerManager player.PlayerManager, gameManager game.GameManager) {
	pm = playerManager
	gm = gameManager

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(playCmd)
	RootCmd.AddCommand(playerCmd)

	playerCmd.AddCommand(createCmd)
	playerCmd.AddCommand(signInCmd)
	playerCmd.AddCommand(signOutCmd)
	playerCmd.AddCommand(retrieveCmd)
	playerCmd.AddCommand(deleteCmd)

	createCmd.Flags().StringP("username", "u", "", "Username. Cannot be empty")
	createCmd.Flags().StringP("password", "p", "", "Password. Cannot be empty")

	signInCmd.Flags().StringP("username", "u", "", "Username. Cannot be empty")
	signInCmd.Flags().StringP("password", "p", "", "Password. Cannot be empty")

	signOutCmd.Flags().StringP("username", "u", "", "Username. Cannot be empty")
	signOutCmd.Flags().StringP("password", "p", "", "Password. Cannot be empty")

	retrieveCmd.Flags().StringP("username", "u", "", "Username. Cannot be empty")

	deleteCmd.Flags().StringP("username", "u", "", "Username. Cannot be empty")
	deleteCmd.Flags().StringP("password", "p", "", "Password. Cannot be empty")

}
