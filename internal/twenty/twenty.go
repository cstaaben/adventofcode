package twenty

import "github.com/spf13/cobra"

var ParentCmd = &cobra.Command{
	Use: "twenty",
}

func init() {
	ParentCmd.AddCommand(day1Cmd)
}
