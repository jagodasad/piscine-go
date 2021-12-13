package piscine

func CountIf(f func(string) bool, tab []string) int {
	ne := 0
	for _, s := range tab {
		if f(s) == true {
			ne++
		}
	}
	return ne
}
