package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for _, v := range arguments[2:] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
