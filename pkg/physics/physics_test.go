package physics

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalPositionOfConstAccelObject(t *testing.T) {
	testCases := []struct {
		x0, v0, a, t float64
		result       float64
	}{
		{0, 0, 26, 5.56, 402.0},
		// Below are test cases from yr2021 day17 example 1
		// {0, 2, -1, 1, 2.0},
		// {0, 2, -1, 2, 3.0},
		// {0, 2, -1, 3, 3.0},
		// {0, 2, -1, 4, 2.0},
		// {0, 2, -1, 5, 0},
		// {0, 2, -1, 6, -3.0},
		// {0, 2, -1, 7, -7.0},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, math.Round(FinalPositionOfConstAccelObject(tc.x0, tc.v0, tc.a, tc.t)))
		})
	}
}

func TestInitialVelocityOfAcceleratingObject(t *testing.T) {
	v, err := InitialVelocityOfConstAccelObject(0, 10, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, v)

	v, err = InitialVelocityOfConstAccelObject(0, 2, 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, v)

	// v, err = InitialVelocityOfConstAccelObject(0, 2, -1, 1)
	// assert.NoError(t, err)
	// assert.Equal(t, 2.0, v)

	// v, err = InitialVelocityOfConstAccelObject(0, -7, -1, 7)
	// assert.NoError(t, err)
	// assert.Equal(t, 2.0, v)
}
