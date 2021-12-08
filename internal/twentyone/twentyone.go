package twentyone

import (
	"github.com/spf13/cobra"

	"github.com/cstaaben/adventofcode/internal/twentyone/day1"
	"github.com/cstaaben/adventofcode/internal/twentyone/day10"
	"github.com/cstaaben/adventofcode/internal/twentyone/day11"
	"github.com/cstaaben/adventofcode/internal/twentyone/day12"
	"github.com/cstaaben/adventofcode/internal/twentyone/day13"
	"github.com/cstaaben/adventofcode/internal/twentyone/day14"
	"github.com/cstaaben/adventofcode/internal/twentyone/day15"
	"github.com/cstaaben/adventofcode/internal/twentyone/day16"
	"github.com/cstaaben/adventofcode/internal/twentyone/day17"
	"github.com/cstaaben/adventofcode/internal/twentyone/day18"
	"github.com/cstaaben/adventofcode/internal/twentyone/day19"
	"github.com/cstaaben/adventofcode/internal/twentyone/day2"
	"github.com/cstaaben/adventofcode/internal/twentyone/day20"
	"github.com/cstaaben/adventofcode/internal/twentyone/day21"
	"github.com/cstaaben/adventofcode/internal/twentyone/day22"
	"github.com/cstaaben/adventofcode/internal/twentyone/day23"
	"github.com/cstaaben/adventofcode/internal/twentyone/day24"
	"github.com/cstaaben/adventofcode/internal/twentyone/day25"
	"github.com/cstaaben/adventofcode/internal/twentyone/day3"
	"github.com/cstaaben/adventofcode/internal/twentyone/day4"
	"github.com/cstaaben/adventofcode/internal/twentyone/day5"
	"github.com/cstaaben/adventofcode/internal/twentyone/day6"
	"github.com/cstaaben/adventofcode/internal/twentyone/day7"
	"github.com/cstaaben/adventofcode/internal/twentyone/day8"
	"github.com/cstaaben/adventofcode/internal/twentyone/day9"
)

var ParentCmd = &cobra.Command{
	Use: "twenty_one",
}

func init() {
	ParentCmd.AddCommand(
		day1.Cmd(),
		day2.Cmd(),
		day3.Cmd(),
		day4.Cmd(),
		day5.Cmd(),
		day6.Cmd(),
		day7.Cmd(),
		day8.Cmd(),
		day9.Cmd(),
		day10.Cmd(),
		day11.Cmd(),
		day12.Cmd(),
		day13.Cmd(),
		day14.Cmd(),
		day15.Cmd(),
		day16.Cmd(),
		day17.Cmd(),
		day18.Cmd(),
		day19.Cmd(),
		day20.Cmd(),
		day21.Cmd(),
		day22.Cmd(),
		day23.Cmd(),
		day24.Cmd(),
		day25.Cmd(),
	)
}
