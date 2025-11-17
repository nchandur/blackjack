package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/nchandur/blackjack/game"
	"github.com/nchandur/blackjack/users"
	"github.com/spf13/cobra"
)

var um users.UserManager
var gm game.GameManager
var GAME_DIR string

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Blackjack CLI game",
	Long: `blackjack is a command-line implementation of the classic card game.
You can play against the dealer using standard Blackjack rules. 
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var print = func(str string) {
			for _, c := range str {
				fmt.Printf("%c", c)
				time.Sleep(20 * time.Millisecond)
			}
			fmt.Println()
		}

		print("======================")
		print(" Welcome to Blackjack ")
		print("======================")
		print("1. Create user\n2. Play Blackjack\n3. Simulate Blackjack\n")

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
	userCmd.AddCommand(resetCmd)

	gameCmd.AddCommand(playCmd)
	gameCmd.AddCommand(simulateCmd)

	createCmd.Flags().StringP("username", "u", "", "Username.")
	createCmd.Flags().StringP("password", "p", "", "Password.")

	createCmd.MarkFlagsRequiredTogether("username", "password")

	signInCmd.Flags().StringP("username", "u", "", "Username.")
	signInCmd.Flags().StringP("password", "p", "", "Password.")

	signInCmd.MarkFlagsRequiredTogether("username", "password")

	signOutCmd.Flags().StringP("username", "u", "", "Username.")
	signOutCmd.Flags().StringP("password", "p", "", "Password.")

	signOutCmd.MarkFlagsRequiredTogether("username", "password")

	retrieveCmd.Flags().StringP("username", "u", "", "Username.")

	retrieveCmd.MarkFlagRequired("username")

	deleteCmd.Flags().StringP("username", "u", "", "Username.")
	deleteCmd.Flags().StringP("password", "p", "", "Password.")

	deleteCmd.MarkFlagsRequiredTogether("username", "password")

	resetCmd.Flags().StringP("username", "u", "", "Username.")
	resetCmd.Flags().StringP("password", "p", "", "Password.")

	resetCmd.MarkFlagsRequiredTogether("username", "password")

	simulateCmd.Flags().StringP("strategy", "s", "basic", "Strategy for the bot to follow")
	simulateCmd.Flags().Int64P("rounds", "r", 10, "Number of rounds to simulate")

}
