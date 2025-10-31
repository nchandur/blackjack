package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Start playing blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to Blackjack!!")
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(playCmd)
}
