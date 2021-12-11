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
	for j := 1; j <= ln; j++ {
		for _, w := range arguments[j] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
