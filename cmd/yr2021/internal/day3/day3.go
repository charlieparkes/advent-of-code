package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charlieparkes/advent-of-code/pkg/files"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:  "03",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	lines, err := load(files.FindPath("cmd/yr2021/internal/day3") + "/input.txt")
	if err != nil {
		return err
	}
	p, err := power(lines) // 1307354
	if err != nil {
		return err
	}
	l, err := lifeSupport(lines)
	if err != nil {
		return err
	}
	log.Default().Printf("pt1: %v, pt2: %v\n", p, l)
	return nil
}

func load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}

func sums(lines []string) []int {
	s := []int{}
	for _, bin := range lines {
		for i, r := range bin {
			if len(s) == i {
				s = append(s, 0)
			}
			s[i] += int(r - '0')
		}
	}
	return s
}

func commonBits(lines []string, prefer int) []int {
	s := sums(lines)
	bits := make([]int, len(s))
	half := len(lines) / 2
	for i, sum := range s {
		if sum > half {
			bits[i] = 1
		} else if sum == half && len(lines)%2 == 0 {
			bits[i] = prefer
		}
	}
	return bits
}

func btoi(bits []int) (int64, error) {
	return strconv.ParseInt(strings.Join(strings.Fields(strings.Trim(fmt.Sprint(bits), "[]")), ""), 2, 64)
}

func flip(bits []int) []int {
	for i := range bits {
		if bits[i] == 1 {
			bits[i] = 0
		} else {
			bits[i] = 1
		}
	}
	return bits
}

func power(lines []string) (int64, error) {
	bits := commonBits(lines, 1)
	g, err := btoi(bits)
	if err != nil {
		return 0, err
	}
	bits = flip(bits)
	e, err := btoi(bits)
	if err != nil {
		return 0, err
	}
	return g * e, nil
}

func oxygen(lines []string) (int64, error) {
	i := 0
	for len(lines) > 1 {
		bits := commonBits(lines, 1)
		reduced := []string{}
		for _, l := range lines {
			if int(l[i]-'0') == bits[i] {
				reduced = append(reduced, l)
			}
		}
		lines = reduced
		i++
	}
	return strconv.ParseInt(lines[0], 2, 64)
}

func co2(lines []string) (int64, error) {
	i := 0
	for len(lines) > 1 {
		bits := flip(commonBits(lines, 1))
		reduced := []string{}
		for _, l := range lines {
			if int(l[i]-'0') == bits[i] {
				reduced = append(reduced, l)
			}
		}
		lines = reduced
		i++
	}
	return strconv.ParseInt(lines[0], 2, 64)
}

func lifeSupport(lines []string) (int64, error) {
	o, err := oxygen(lines)
	if err != nil {
		return 0, err
	}
	c, err := co2(lines)
	if err != nil {
		return 0, err
	}
	return c * o, nil
}
