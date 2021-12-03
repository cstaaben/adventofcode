package twentyone

import "github.com/spf13/cobra"

var ParentCmd = &cobra.Command{
	Use: "twenty_one",
}

func init() {
	ParentCmd.AddCommand(
		day1Cmd,
		day2Cmd,
		day3Cmd,
		day4Cmd,
		day5Cmd,
		day6Cmd,
		day7Cmd,
		day8Cmd,
		day9Cmd,
		day10Cmd,
		day11Cmd,
		day12Cmd,
		day13Cmd,
		day14Cmd,
		day15Cmd,
		day16Cmd,
		day17Cmd,
		day18Cmd,
		day19Cmd,
		day20Cmd,
		day21Cmd,
		day22Cmd,
		day23Cmd,
		day24Cmd,
		day25Cmd,
	)
}
