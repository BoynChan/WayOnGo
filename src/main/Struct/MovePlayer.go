package main

import "math"

type Vector struct {
	X, Y float32
}

func (v *Vector) Add(other Vector) {
	v.Y += other.Y
	v.X += other.X
}

func (v *Vector) Sub(other Vector) {
	v.Y -= other.Y
	v.X -= other.X
}

func (v Vector) DistanceTo(other Vector) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func main() {

}
