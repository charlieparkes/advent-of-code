package vector

import "math"

type Vector struct {
	X, Y float64
}

func (v Vector) Equal(ov Vector) bool {
	const epsilon = 1e-16
	return math.Abs(v.X-ov.X) < epsilon && math.Abs(v.Y-ov.Y) < epsilon
}

func (v Vector) Add(ov Vector) Vector { return Vector{v.X + ov.X, v.Y + ov.Y} }

func (v Vector) Sub(ov Vector) Vector { return Vector{v.X - ov.X, v.Y - ov.Y} }

func (v Vector) Mul(m float64) Vector { return Vector{m * v.X, m * v.Y} }

func (v Vector) Dot(ov Vector) float64 { return v.X*ov.X + v.Y*ov.Y }

func (v Vector) Norm() float64 { return math.Sqrt(v.Dot(v)) }

func (v Vector) Distance(ov Vector) float64 { return v.Sub(ov).Norm() }
