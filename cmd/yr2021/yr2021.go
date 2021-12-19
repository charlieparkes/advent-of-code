package yr2021

import (
	"github.com/charlieparkes/advent-of-code/cmd/yr2021/internal/day17"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "2021",
}

func init() {
	Cmd.AddCommand(
		day17.Cmd,
	)
}
