package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "Start playing blackjack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Blackjack!!")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
