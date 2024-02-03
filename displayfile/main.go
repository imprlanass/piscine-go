package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	minBuffer := make([]byte, 1)
	for {
		_, err := file.Read(minBuffer)
		if err != nil {
			break
		}
		fmt.Print(string(minBuffer))
	}
}
