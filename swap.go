package piscine

func Swap(a *int, b *int) {
	pA := *a
	*a = *b
	*b = pA
}
