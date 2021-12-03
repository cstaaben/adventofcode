package cmd

import "github.com/cstaaben/adventofcode/internal/twentyone"

func init() {
	rootCmd.AddCommand(twentyone.ParentCmd)
}
