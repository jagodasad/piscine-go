package piscine

func IsNumeric(s string) bool {
	characters := []rune(s)
	for i := range s {
		if '0' <= characters[i] && '9' >= characters[i] {
		} else {
			return false
		}
	}
	return true
}
