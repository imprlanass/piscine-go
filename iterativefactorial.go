package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb < 0 || nb >= 66 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		res *= int(i)
	}
	return res
}
