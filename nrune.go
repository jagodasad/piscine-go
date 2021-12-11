package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n <= len(a) && n >= 1 {
		return a[n-1]
	}
	return 0
}
