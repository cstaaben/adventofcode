package cmd

import "github.com/cstaaben/adventofcode/internal/twenty"

func init() {
	rootCmd.AddCommand(twenty.ParentCmd)
}
