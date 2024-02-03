package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			w := y + 1
			for z := x; z <= '9'; z++ {
				for ; w <= '9'; w++ {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(' ')
					z01.PrintRune(z)
					z01.PrintRune(w)
					if x < '9' || y < '8' || z < '9' || w < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				w = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
