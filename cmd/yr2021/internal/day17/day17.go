package day17

import (
	"fmt"
	"log"

	"github.com/charlieparkes/advent-of-code/pkg/vector"
	"github.com/spf13/cobra"
)

// target area: x=175..227, y=-134..-79
var TARGET = Target{
	Point{175, -134},
	Point{227, -79},
}

// target area: x=20..30, y=-10..-5
// var TARGET = Target{
// 	Point{20, -10},
// 	Point{30, -5},
// }

var Cmd = &cobra.Command{
	Use:  "17",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	var v vector.Vector
	err := solve(TARGET, func(solution vector.Vector) bool {
		v = solution
		return false
	})
	if err != nil {
		return err
	}
	height := 0.0
	for vy := v.Y; vy > 0; vy-- {
		height += vy
	}
	n := 0
	_ = solve(TARGET, func(_ vector.Vector) bool {
		n++
		return true
	})
	log.Default().Printf("%+v, height: %v, total: %v", v, height, n)
	return nil
}

type Point struct {
	X, Y int
}

type Target struct {
	P1, P2 Point
}

func (t *Target) Contains(p vector.Vector) bool {
	return float64(t.P1.X) <= p.X && p.X <= float64(t.P2.X) && float64(t.P1.Y) <= p.Y && p.Y <= float64(t.P2.Y)
}

// return value of callback determines should continue after finding a solution
func solve(t Target, callback func(vector.Vector) bool) error {
	minX := func(bound int) int {
		vel := 0
		for pos := 0; pos < bound; pos += vel {
			vel++
		}
		return vel
	}
	for y := -t.P1.Y - 1; y >= t.P1.Y; y-- {
		for x := minX(t.P1.X); x <= t.P2.X; x++ {
			v := vector.Vector{X: float64(x), Y: float64(y)}
			if err := check(t, v); err == nil && !callback(v) {
				return nil
			}
		}
	}
	return fmt.Errorf("no solution found for target: %+v", t)
}

func check(t Target, solution vector.Vector) error {
	loc := vector.Vector{X: 0, Y: 0}
	if t.Contains(loc) {
		return nil
	}
	v := solution
	for loc.X <= float64(t.P2.X) && loc.Y >= float64(t.P1.Y) {
		loc = loc.Add(v)
		if t.Contains(loc) {
			return nil
		}
		if v.X > 0 {
			v = v.Sub(vector.Vector{X: 1, Y: 0})
		} else if v.X < 0 {
			v = v.Add(vector.Vector{X: 1, Y: 0})
		}
		v = v.Sub(vector.Vector{X: 0, Y: 1})
	}
	return fmt.Errorf("invalid solution: %+v", solution)
}
