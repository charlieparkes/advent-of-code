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
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, math.Round(FinalPositionOfConstAccelObject(tc.x0, tc.v0, tc.a, tc.t)))
		})
	}
}

func TestInitialVelocityOfAcceleratingObject(t *testing.T) {
	v, err := InitialVelocityOfConstAccelObject(0, 402, 26, 5.56)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, math.Round(v))
}
