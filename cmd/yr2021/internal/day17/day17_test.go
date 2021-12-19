package day17

import (
	"testing"

	"github.com/charlieparkes/advent-of-code/pkg/vector"
	"github.com/stretchr/testify/assert"
)

func TestTargetContains(t *testing.T) {
	t1 := Target{Point{X: 1, Y: 1}, Point{X: 1, Y: 1}}
	assert.False(t, t1.Contains(vector.Vector{X: 0, Y: 1}))
	assert.True(t, t1.Contains(vector.Vector{X: 1, Y: 1}))
	assert.False(t, t1.Contains(vector.Vector{X: 1, Y: 2}))

	t2 := Target{Point{X: 15, Y: 0}, Point{X: 21, Y: 10}}
	assert.True(t, t2.Contains(vector.Vector{X: 18, Y: 8}))
	assert.False(t, t2.Contains(vector.Vector{X: 13, Y: 5}))
}

func TestSolve(t *testing.T) {
	var v vector.Vector
	err := solve(Target{
		Point{20, -10},
		Point{30, -5},
	}, func(solution vector.Vector) bool {
		v = solution
		return false
	})
	assert.NoError(t, err)
	assert.Equal(t, vector.Vector{X: 6, Y: 9}, v)

	solutions := []vector.Vector{}
	_ = solve(Target{
		Point{20, -10},
		Point{30, -5},
	}, func(solution vector.Vector) bool {
		solutions = append(solutions, solution)
		return true
	})
	assert.Equal(t, 112, len(solutions))
}

func TestCheck(t *testing.T) {
	t1 := Target{Point{X: 20, Y: -10}, Point{X: 30, Y: -5}}
	assert.NoError(t, check(t1, vector.Vector{X: 7, Y: 2}))
	assert.NoError(t, check(t1, vector.Vector{X: 6, Y: 3}))
	assert.NoError(t, check(t1, vector.Vector{X: 9, Y: 0}))
	assert.Error(t, check(t1, vector.Vector{X: 17, Y: -4}))
	assert.NoError(t, check(t1, vector.Vector{X: 6, Y: 9}))
}
