package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User Management",
	Long: `Manage Blackjack users and their game data.

Use this command to create, sign in, sign out, view, or delete users.
You must be signed in to play and record Blackjack stats.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("This is the User Management Section")

		return nil
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user with a username and password.",
	Long: `Create a new user account.

Each user has a username and password. All game stats and Blackjack rounds 
played while signed in with this account will be stored under it.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = um.Create(username, password)

		if err != nil {
			return err
		}

		err = um.SignIn(username, password)

		fmt.Printf("%s has been successfully created!\n", username)
		return nil
	},
}

var signInCmd = &cobra.Command{
	Use:   "signin",
	Short: "Sign in with your username and password.",
	Long: `Sign in to your Blackjack account.

You must be signed in before playing so that your game history and 
statistics are saved. Signing in automatically signs out any active user.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = um.SignIn(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully signed in\n", username)
		return nil
	},
}

var signOutCmd = &cobra.Command{
	Use:   "signout",
	Short: "Sign out of the current user session.",
	Long: `Sign out of your current Blackjack session.

Once signed out, you can either sign in again or let another user sign in.
Only one user can be signed in at a time.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = um.SignOut(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully signed out\n", username)
		return nil
	},
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve statistics for a specific user.",
	Long: `View the recorded Blackjack statistics for any user.

You can retrieve stats such as games played, wins, losses, and balance 
for a given username.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		play, err := um.Retrieve(username)

		if err != nil {
			return err
		}

		fmt.Println(play.String())

		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user account.",
	Long: `Permanently remove a user and their stored game data.

Once deleted, all game history and statistics for that user are lost.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = um.Delete(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully deleted\n", username)
		return nil
	},
}
