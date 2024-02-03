package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	s := make([]int, max-min)
	for i := range s {
		s[i] = i + min
	}
	return s
}
