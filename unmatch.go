package piscine

func Unmatch(a []int) int {
for i := 0; i < Lent4(a); i++ {
		for j := i + 1; j < Lent4(a); j++ {
			if a[i] == a[j] {
				a[i] = -1
				a[j] = -1
				break
			}
		}
	}
	for k := 0; k < Lent4(a); k++ {
		if a[k] != -1 {

			return a[k]
		}
	}
return -1
}

func Lent4(d []int) int{
	inc:=0
	for _, _= range d{
		inc ++
	}
	 return inc
}