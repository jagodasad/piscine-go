package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		works := true

		_, err1 := strconv.Atoi(args[0])
		_, err2 := strconv.Atoi(args[2])

		if err1 != nil || err2 != nil {
			works = false
		}

		if works {
			os.Stdout.WriteString(Maths(args))
		}
	}
}

func Maths(s []string) string {
	nb1, _ := strconv.Atoi(s[0])
	nb2, _ := strconv.Atoi(s[2])
	var result int
	var output string

	switch s[1] {
	case "+":
		if nb1 == 9223372036854775807 || nb2 == 9223372036854775807 {
			return ""
		}
		result = nb1 + nb2
		output = changeToString(result)
	case "-":
		if nb1 == 9223372036854775807 || nb2 == 9223372036854775807 {
			return ""
		}
		result = nb1 - nb2
		output = changeToString(result)
	case "/":
		if nb2 == 0 || nb1 == 0 {
			return "No division by 0\n"
		}
		result = nb1 / nb2
		output = changeToString(result)
	case "%":
		if nb2 == 0 || nb1 == 0 {
			return "No modulo by 0\n"
		}
		result = nb1 % nb2
		output = changeToString(result)
	case "*":
		if nb1 == 9223372036854775807 || nb2 == 9223372036854775807 {
			return ""
		}
		result = nb1 * nb2
		output = changeToString(result)
	}
	return output
}

func changeToString(nbr int) string {
	var output string
	output = strconv.Itoa(nbr)
	return output + "\n"
}
