package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/nchandur/blackjack/simulator"
	"github.com/spf13/cobra"
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Simulate rounds of Blackjack",
	RunE: func(cmd *cobra.Command, args []string) error {

		strategy, err := cmd.Flags().GetString("strategy")

		if err != nil {
			return fmt.Errorf("failed to simulate game: %v", err)
		}

		stratPath := ""

		switch strategy {
		case "basic":
			stratPath = "strategy/basic.bin"
		case "h17":
			stratPath = "strategy/h17.bin"
		case "s17":
			stratPath = "strategy/s17.bin"
		case "enhc":
			stratPath = "strategy/enhc.bin"
		default:
			stratPath = "strategy/basic.bin"
		}

		sim, err := simulator.NewSimulator(filepath.Join(GAME_DIR, stratPath))

		if err != nil {
			return fmt.Errorf("failed to simulate game: %v", err)
		}

		rounds, err := cmd.Flags().GetInt64("rounds")

		if err != nil {
			return fmt.Errorf("failed to simulate game: %v", err)
		}

		sim.Simulate(uint64(rounds))

		fmt.Println(sim.Report())

		return nil

	},
}
