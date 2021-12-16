package piscine

func Join(strs []string, sep string) string {
	strV := ""

	for i := 0; i < Con(strs); i++ {
		strV += strs[i]
		if i < Con(strs)-1 {
			strV += sep
		}
	}
	return strV
}

func Con(d []string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}
