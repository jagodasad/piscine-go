package piscine

func ToLower(s string) string {
	characters := []rune(s)

	for i, c := range s {
		if 65 <= c && 90 >= c {
			characters[i] = rune(c + 32)
		}
	}
	return string(characters)
}
