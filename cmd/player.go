package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Player Management",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("This is the Player Management Section")

		return nil
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new player",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("creates new player")
		return nil
	},
}

var signInCmd = &cobra.Command{
	Use:   "signin",
	Short: "Sign in new player",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("signin new player")
		return nil
	},
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve existing player",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("retrieve existing player")
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete existing player",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("delete existing player")
		return nil
	},
}
