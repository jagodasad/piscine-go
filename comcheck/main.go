package main

import (
	"fmt"
	"os"
)

// func main() {
// 	r := ""
// 	arg := os.Args[1:]

// 	check := []string{"01", "galaxy", "galaxy 01"}
// 	for _, res := range arg {
// 		for _, i := range check {
// 			if res == i {
// 				r += "Alert!!!\n"
// 				break
// 			}
// 		}
// 	}
// 	fmt.Print(r)
// }
func main() {
	r := 0
	arg := os.Args[1:]
	check := []string{"01", "galaxy", "galaxy 01"}
	for i := range arg {
		for _, s := range check {
			if arg[i] == s {
				r++
			}
		}
	}
	if r >= 1 {
		fmt.Println("Alert!!!")
	}
}
