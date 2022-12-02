package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "advent",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		return viper.BindPFlags(cmd.Flags())
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug logging")

	rootCmd.AddCommand(newAddCommand())
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
