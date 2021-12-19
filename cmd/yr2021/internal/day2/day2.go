package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charlieparkes/advent-of-code/pkg/files"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:  "02",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	path := files.FindPath("cmd/yr2021/internal/day2") + "/input.txt"
	ans1, err := pt1(path)
	if err != nil {
		return err
	}
	ans2, err := pt2(path)
	if err != nil {
		return err
	}
	log.Default().Printf("pt1: %v, pt2: %v", ans1, ans2)
	return nil
}

func pt1(path string) (int, error) {
	var x, y int
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		cmd := s.Text()
		parts := strings.SplitN(cmd, " ", 2)
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		switch parts[0][0] {
		case 'f':
			x += i
		case 'u':
			y += i
		case 'd':
			y -= i
		}
	}
	return x * -y, nil
}

func pt2(path string) (int, error) {
	var aim, x, y int
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		cmd := s.Text()
		parts := strings.SplitN(cmd, " ", 2)
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		switch parts[0][0] {
		case 'f':
			x += i
			y += aim * i
		case 'u':
			aim += i
		case 'd':
			aim -= i
		}
	}
	return x * -y, nil
}
