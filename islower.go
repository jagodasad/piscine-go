package piscine

func IsLower(s string) bool {
	characters := []rune(s)
	count := 0
	for i := range s {
		if 'a' <= characters[i] && 'z' >= characters[i] {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
