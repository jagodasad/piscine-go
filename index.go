package piscine

func Index(s string, toFind string) int {
	sRune := []rune(s)
	toFindRune := []rune(toFind)
	counter := 0

	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if sRune[i] == toFindRune[0] {
			counter = counter + (i + 1)
			break
		}
	}
	if counter == 0 {
		return -1
	}
	return counter - 1
}
