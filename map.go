package piscine

func Map(f func(int) bool, a []int) []bool {
	m := make(map[int]bool)
	result := make([]bool, len(a))

	var resultIndex int

	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}

		if f(v) {
			result[resultIndex] = true
		} else {
			result[resultIndex] = false
		}

		m[v] = f(v)
		resultIndex++
	}

	return result[:resultIndex]
}
