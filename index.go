package piscine

func Index(s string, toFind string) int {
	sRune := []rune(s)
	toFindRune := []rune(toFind)
	counter := 0

	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if sRune[i] == toFindRune[0] && len(toFind) <= 1 {
			counter = counter + (i + 1)
			break
		}
		if len(toFind) > 1 && sRune[i] == toFindRune[0] && sRune[i+1] == toFindRune[1] {
			counter = counter + (i + 1)
			break
		}
		//if len(toFind) > 1 && sRune[i+1] != toFindRune[1] {
		//	return -1
		//}
	}

	if counter == 0 {
		return -1
	}

	return counter - 1
}
