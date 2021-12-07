package piscine

func AlphaCount(s string) int {
	// sentence := []rune(s)
	counter := 0
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			counter++
		}
	}
	return counter
}
