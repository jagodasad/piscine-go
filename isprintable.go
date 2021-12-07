package piscine

func IsPrintable(s string) bool {
	characters := []rune(s)
	for i := range s {
		if 32 <= characters[i] && 127 >= characters[i] {
		} else {
			return false
		}
	}
	return true
}
