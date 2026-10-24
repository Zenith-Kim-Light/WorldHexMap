package util

func CalcNearestPowersOfTwo(width, height int32) (nearestWidth, nearestHeight int32) {
	if width <= 0 {
		nearestWidth = 1
	} else {
		w := width - 1
		w |= w >> 1
		w |= w >> 2
		w |= w >> 4
		w |= w >> 8
		w |= w >> 16
		nearestWidth = w + 1
	}

	if height <= 0 {
		nearestHeight = 1
	} else {
		h := height - 1
		h |= h >> 1
		h |= h >> 2
		h |= h >> 4
		h |= h >> 8
		h |= h >> 16
		nearestHeight = h + 1
	}
	return nearestWidth, nearestHeight
}
