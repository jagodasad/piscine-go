package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg1 := os.Args
	arg2 := os.Args
	ln := 0
	for i := range arg1 {
		ln = i
	}
	for i := 1; i <= ln; i++ {
		for j := 0; j <= ln; j++ {
			if arg1[i] < arg2[j] {
				arg2[j], arg1[i] = arg1[i], arg2[j]
			}
		}
	}
	for i := 0; i <= ln; i++ {
		for _, w := range arg2[i] {
			z01.PrintRune(w)
		}

		z01.PrintRune(10)
	}
}
