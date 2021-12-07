package piscine

func IsAlpha(s string) bool {
	/*	counter := 0
		for _, c := range s {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')  || (c >= '0' && c <= '9') {
				return true
			}
		} else {
		return false
		}*/
	characters := []rune(s)
	count := 0
	for i := range s {
		if 'a' <= characters[i] && 'z' >= characters[i] || 'A' <= characters[i] && 'Z' >= characters[i] || '0' <= characters[i] && '9' >= characters[i] {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
