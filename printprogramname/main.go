package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	abs := arguments[0]

	var result string

	for i := len(abs) - 1; i >= 0; i-- {
		if string(abs[i]) == "/" {
			break
		} else {
			result += string(abs[i])
		}
	}
	s := piscine.StrRev(result)
	for _, l := range s {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
