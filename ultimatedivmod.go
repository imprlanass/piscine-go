package piscine

func UltimateDivMod(a *int, b *int) {
	pA := *a / *b
	*b = *a % *b
	*a = pA
}
