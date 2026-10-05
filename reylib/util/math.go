package util

import "math"

func CalcMagnitude(x float64, y float64) float64 {
	sum := x + y
	pow := math.Pow(sum, 2)
	return math.Sqrt(pow)
}
