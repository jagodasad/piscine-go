package piscine

func BasicJoin(elems []string) string {
	myStr := ""
	for _, w := range elems {
		myStr = myStr + w
	}
	return myStr
}
