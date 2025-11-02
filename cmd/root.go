package cmd

import (
	"fmt"
	"log"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/users"
	"github.com/spf13/cobra"
)

var um users.UserManager
var gm game.GameManager

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Start playing blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to Blackjack!!")
		return nil
	},
}

func Execute(userManager users.UserManager, gameManager game.GameManager) {
	um = userManager
	gm = gameManager

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(gameCmd)
	RootCmd.AddCommand(userCmd)

	userCmd.AddCommand(createCmd)
	userCmd.AddCommand(signInCmd)
	userCmd.AddCommand(signOutCmd)
	userCmd.AddCommand(retrieveCmd)
	userCmd.AddCommand(deleteCmd)

	gameCmd.AddCommand(playCmd)

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
