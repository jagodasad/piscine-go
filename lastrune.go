package piscine

func LastRune(s string) rune {
	sentence := []rune(s)
	i := len(sentence)
	return sentence[i-1]
}
