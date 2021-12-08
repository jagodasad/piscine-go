package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	ln := 0
	for i := range arguments {
		ln = i
	}
	for i := 1; i < ln; i++ {
		for _, w := range arguments[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
