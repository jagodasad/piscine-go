package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				if k == i {
					for l := j + 1; l <= '9'; l++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						if i != '9' || j != '8' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				} else {
					for l := '0'; l <= '9'; l++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
