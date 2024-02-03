package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		tab = append(tab, n%10)
		n = n / 10
	}
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	for _, v := range tab {
		z01.PrintRune(rune(v + '0'))
	}
}
