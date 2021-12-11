package piscine

func ToUpper(s string) string {
	characters := []rune(s)

	for i, c := range s {
		if 97 <= c && 122 >= c {
			characters[i] = rune(c - 32)
		}
	}
	return string(characters)
}
