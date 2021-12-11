package piscine

func IsUpper(s string) bool {
	characters := []rune(s)
	count := 0
	for i := range s {
		if 'A' <= characters[i] && 'Z' >= characters[i] {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
