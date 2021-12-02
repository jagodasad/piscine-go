package piscine

func BasicAtoi(s string) int {
	if string(s[0]) == "-" {
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}
}
