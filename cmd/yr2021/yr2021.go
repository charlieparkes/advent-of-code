package yr2021

import (
	"github.com/charlieparkes/advent-of-code/cmd/yr2021/internal/day1"
	"github.com/charlieparkes/advent-of-code/cmd/yr2021/internal/day17"
	"github.com/charlieparkes/advent-of-code/cmd/yr2021/internal/day2"
	"github.com/charlieparkes/advent-of-code/cmd/yr2021/internal/day3"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "2021",
}

func init() {
	Cmd.AddCommand(
		day1.Cmd,
		day2.Cmd,
		day3.Cmd,
		day17.Cmd,
	)
}
