package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if nb <= 65 && power > 0 {
		for i := 0; i < power; i++ {
			result = result * nb
		}
	} else if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	return result
}
