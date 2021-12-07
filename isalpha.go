package piscine

func IsAlpha(s string) bool {
	characters := []rune(s)
	for i := range s {
		if 'a' <= characters[i] && 'z' >= characters[i] || 'A' <= characters[i] && 'Z' >= characters[i] || '0' <= characters[i] && '9' >= characters[i] {
		} else {
			return false
		}
	}
	return true
}
