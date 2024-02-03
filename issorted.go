package piscine

func CheckOrder(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	length := 0
	for i := range a {
		length = i + 1
	}

	asc := true
	desc := true

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			desc = false
		}
	}

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asc = false
		}
	}

	return asc || desc
}
