package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, v := range arguments {
		for _, vtwo := range v {
			z01.PrintRune(vtwo)
		}
		z01.PrintRune('\n')
	}
}
