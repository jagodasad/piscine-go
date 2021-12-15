package main

import (
	"fmt"
	"os"
)

func main() {
	r := ""
	arg := os.Args[1:]

	check := []string{"01", "galaxy", "galaxy 01"}
	for _, res := range arg {
		for _, i := range check {
			if res == i {
				r += "Alert!!!\n"
				break
			}
		}
	}
	fmt.Print(r)
}
