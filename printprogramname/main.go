package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, w := range arguments[0] {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
