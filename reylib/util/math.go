package util

import "math"

func CalcNearestPowersOfTwo(width, height int32) (nearestWidth, nearestHeight int32) {
	floatWidth := float64(width)
	floatHeight := float64(height)

	for i := 0.0; ; i++ {
		powerOfTow := math.Pow(2, i)
		if 0.0 == nearestWidth && powerOfTow >= floatWidth {
			nearestWidth = int32(powerOfTow)
		}
		if 0.0 == nearestHeight && powerOfTow >= floatHeight {
			nearestHeight = int32(powerOfTow)
		}
		if nearestWidth != 0 && nearestHeight != 0 {
			break
		}
	}
	return
}
