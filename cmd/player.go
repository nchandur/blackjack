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

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = pm.Create(username, password)

		if err != nil {
			return err
		}

		err = pm.SignIn(username, password)

		fmt.Printf("%s has been successfully created!\n", username)
		return nil
	},
}

var signInCmd = &cobra.Command{
	Use:   "signin",
	Short: "Sign in new player",
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = pm.SignIn(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully signed in\n", username)
		return nil
	},
}

var signOutCmd = &cobra.Command{
	Use:   "signout",
	Short: "Sign out",
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = pm.SignOut(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully signed out\n", username)
		return nil
	},
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve existing player",
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		play, err := pm.Retrieve(username)

		if err != nil {
			return err
		}

		fmt.Println(play.String())

		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete existing player",
	RunE: func(cmd *cobra.Command, args []string) error {

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		err = pm.Delete(username, password)

		if err != nil {
			return err
		}

		fmt.Printf("%s has been successfully deleted\n", username)
		return nil
	},
}
