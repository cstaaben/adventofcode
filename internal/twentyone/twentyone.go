package twentyone

import "github.com/spf13/cobra"

var ParentCmd = &cobra.Command{
	Use: "twenty_one",
}

func init() {
	ParentCmd.AddCommand(day1Cmd, dayTwoCmd)
}
