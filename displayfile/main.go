package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("File name missing")
	} else if len(arg) > 2 {
		fmt.Println("Too many arguments")
	} else if len(arg) == 2 {
		data, err := ioutil.ReadFile(arg[1])

		if err != nil {
			fmt.Println(err.Error())
			
		} else {
		fmt.Println(string(data))
		}
	}
}