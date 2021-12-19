package physics

import (
	"errors"
	"math"
)

// vf^2 = v0^2 + 2ax
// func FinalVelocityOfConstAccelObjectGivenPosition(v0, a, x)
// sqrt(v0^2+2ax)

// vf = v0 + at
// func FinalVelocityOfConstAccelObjectGivenTime(v0, a, t)

// xf = x0 + v0*t + 0.5at^2
func FinalPositionOfConstAccelObject(x0, v0, a, t float64) float64 {
	return x0 + v0*t + 0.5*a*math.Pow(t, 2)
}

// v0 = -(2(x0) - 2(xf) + at^2) / 2t
func InitialVelocityOfConstAccelObject(x0, xf, a, t float64) (float64, error) {
	if t == 0 {
		return 0, errors.New("t may not be zero")
	}
	return -((2*x0 - 2*xf + a*math.Pow(t, 2)) / (2 * t)), nil
}
