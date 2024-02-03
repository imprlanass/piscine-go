package piscine

func CollatzCountdown(start int) int {
	steps := 0
	if start <= 0 {
		return -1
	}
	for start > 0 {
		if start%2 == 0 {
			start /= 2
		} else if start == 1 {
			break
		} else if start%2 != 0 {
			start = (start * 3) + 1
		}
		steps++
	}
	return steps
}
