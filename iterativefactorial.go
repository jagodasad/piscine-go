package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb <= 25 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	} else if nb < 0 {
		return 0
	} else {
		return 0
	}

	return result
}
