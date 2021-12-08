package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cstaaben/adventofcode/cmd/add"
)

var rootCmd = &cobra.Command{
	Use: "advent",
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringP("input_file", "i", "", "Input file for puzzle")
	rootCmd.PersistentFlags().Bool("part_one", false, "Run part one of the day's puzzle")
	rootCmd.PersistentFlags().Bool("part_two", false, "Run part two of the day's puzzle")
	rootCmd.PersistentFlags().BoolP("all", "a", false, "Run all parts of the day's puzzle")
	cobra.CheckErr(viper.BindPFlags(rootCmd.PersistentFlags()))

	rootCmd.AddCommand(add.Cmd())
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
