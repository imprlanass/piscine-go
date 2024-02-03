package piscine

import "github.com/01-edu/z01"

func PrintBase(n int, base string) {
	b := len(base)
	i := 0
	if n == 0 {
		z01.PrintRune('0')
	}
	for j := 1; j <= n%b; j++ {
		i++
	}
	for j := -1; j >= n%b; j-- {
		i++
	}
	if n/b != 0 {
		PrintBase(n/b, base)
	}
	z01.PrintRune(rune(base[i]))
	return
}

func PrintNbrBase(nbr int, base string) {
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if i != j && base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
	}
	PrintBase(nbr, base)
}
