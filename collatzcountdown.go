package piscine

func CollatzCountdown(start int) int {
	step := 0
	if start > 0 {
		for start > 1 {
			if start%2 == 1 {
				start = (3 * start) + 1
			} else {
				start /= 2
			}
			step++
		}
	} else {
		return -1
	}
	return step
}
