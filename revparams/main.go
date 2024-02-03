package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	indexOfArgs := len(arguments) - 1
	for indexOfArgs >= 0 {
		for _, v := range arguments[indexOfArgs] {
			z01.PrintRune(v)
		}
		indexOfArgs--
		z01.PrintRune('\n')
	}
}
