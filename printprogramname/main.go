package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	abs := arguments[0]
	for i, l := range abs {
		if i > 1 {
			z01.PrintRune(rune(l))
		}
	}
	z01.PrintRune('\n')
}
