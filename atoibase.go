package piscine

func AtoiBase(s string, base string) int {
	for i := 0; i < len(base); i++ {
		for j := 0; j < i; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	if len(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		return 0
	}
	res := 0

	for i := 0; i < len(s); i++ {
		j := 0
		for ; j < len(base); j++ {
			if s[i] == base[j] {
				break
			}
		}
		res = res + j*IterativePower(len(base), len(s)-i-1)

	}
	return res
}
