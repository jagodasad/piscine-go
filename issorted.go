package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	a := true
	b := true
	for i := 1; i < Lent4(a); i++ {
		if !(f(a[i-1], a[i]) <= 0 ) {
			a = false
		}
	}
	for i :=1; i < Lent4(a); i++ {
		if !(f(a[i-1], a[i]) <= 0 ) {
			b = false
		}
	}
	return a || b

}

func F(a, b int) int {
	if a > b {
		return 1
	} else if a==b {
		return 0
	} else {
		return -1
	}
}
func Lent4(d []int) int {
	inc := 0
	for _, _ = range d{
		inc++
	}
	return inc
}