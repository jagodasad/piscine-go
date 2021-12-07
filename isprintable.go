package piscine

func IsPrintable(s string) bool {
	characters := []rune(s)
	for i := range s {
		if 33 <= characters[i] && 126 >= characters[i] {
		} else {
			return false
		}
	}
	return true
}
