package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if nb >= 0 && nb <= 65 {
		for i := 0; i < power; i++ {
			result = result * nb
		}
	} else if nb < 0 {
		return 0
	} else {
		return 0
	}
	return result
}
