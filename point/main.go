package main

import (
	"github.com/01-edu/z01"
)

var output string = "x = 42, y = 21\n"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	for _, r := range output {
		z01.PrintRune(r)
	}
}
