package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if result <= 25 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	} else {
		return 0
	}
	return result
}
