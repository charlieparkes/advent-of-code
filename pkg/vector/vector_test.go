package vector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		v1, v2 Vector
		result bool
	}{
		{
			Vector{X: 1, Y: 1},
			Vector{X: 1, Y: 1},
			true,
		},
		{
			Vector{X: 1.234, Y: 1.234},
			Vector{X: 1.234, Y: 1.234},
			true,
		},
		{
			Vector{X: 1, Y: 1},
			Vector{X: 2, Y: 2},
			false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, tc.v1.Equal(tc.v2))
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		v1, v2, result Vector
	}{
		{
			Vector{X: 1, Y: 1},
			Vector{X: 1, Y: 1},
			Vector{X: 2, Y: 2},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, tc.v1.Add(tc.v2))
		})
	}
}

func TestSub(t *testing.T) {
	testCases := []struct {
		v1, v2, result Vector
	}{
		{
			Vector{X: 1, Y: 1},
			Vector{X: 1, Y: 1},
			Vector{X: 0, Y: 0},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, tc.v1.Sub(tc.v2))
		})
	}
}

func TestMul(t *testing.T) {
	testCases := []struct {
		v      Vector
		m      float64
		result Vector
	}{
		{
			Vector{X: 2, Y: 1},
			3.0,
			Vector{X: 6, Y: 3},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, tc.v.Mul(tc.m))
		})
	}
}

func TestDot(t *testing.T) {
	testCases := []struct {
		v1, v2 Vector
		result float64
	}{
		{
			Vector{X: 2, Y: 2},
			Vector{X: 3, Y: 3},
			12.0,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, tc.v1.Dot(tc.v2))
		})
	}
}
