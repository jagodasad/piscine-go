package main

import (
	"os"

	"github.com/01-edu/z01"
)

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

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
	s := StrRev(result)
	for _, l := range s {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
