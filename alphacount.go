package piscine

func AlphaCount(s string) int {
	// sentence := []rune(s)
	counter := 0
	for _, c := range s {
		n := c
		for i := 'a'; i <= 'z'; i++ {
			if n == i {
				counter++
			}
			for i := 'A'; i <= 'Z'; i++ {
				if n == i {
					counter++
				}
			}
		}
	}
	return counter
}
