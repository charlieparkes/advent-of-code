package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charlieparkes/advent-of-code/cmd/yr2021"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "adventofcode",
}

func main() {
	RootCmd.AddCommand(
		yr2021.Cmd,
	)
	yr2021.Cmd.AddCommand(&cobra.Command{
		Use: "all",
		RunE: func(_ *cobra.Command, args []string) error {
			for _, c := range yr2021.Cmd.Commands() {
				if c.Use == "all" {
					continue
				}
				args := []string{yr2021.Cmd.Use, c.Use}
				RootCmd.SetArgs(args)
				log.Default().Println("run", args)
				if err := RootCmd.Execute(); err != nil {
					return err
				}
			}
			return nil
		},
	})
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
