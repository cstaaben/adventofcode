// THIS FILE IS AUTOGENERATED; DO NOT EDIT

package cmd

import "github.com/cstaaben/adventofcode/internal/days"

func init() {
	rootCmd.AddCommand(
		days.NewFiveCommand(),
		days.NewFourCommand(),
		days.NewOneCommand(),
		days.NewThreeCommand(),
		days.NewTwoCommand(),
		days.NewSixCommand(),
	)
}
