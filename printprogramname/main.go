package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	abs := arguments[0][33:]
	for _, x := range abs {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
