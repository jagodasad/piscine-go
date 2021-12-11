package piscine

func UltimateDivMod(a *int, b *int) {
	c := 0
	// var pointer *int = &c
	c = *a / *b
	*b = *a % *b
	*a = c
}
