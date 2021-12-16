package piscine

func Max(a []int) int {
	lem := 0
	for l := range a {
		lem = l + 1
	}
	i := 1
	for i < lem {
		if a[i-1] > a[i] {
			temp := i
			a[i] = a[i-1]
			a[i-1] = a[temp]
			i = 1
		} else {
			i++
		}
	}
	return a[lem-1]
}
