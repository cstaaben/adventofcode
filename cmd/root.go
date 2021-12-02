package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cstaaben/adventofcode/internal/twentyone"
)

var rootCmd = &cobra.Command{
	Use: "advent",
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug logging")
	cobra.CheckErr(viper.BindPFlags(rootCmd.PersistentFlags()))

	rootCmd.AddCommand(twentyone.DayOneCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
