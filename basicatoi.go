package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, v := range s {
		if v >= '1' || v <= '9' {
			res = res*10 + int(v) - '0'
		} else {
			return 0
		}
	}
	return res
}
