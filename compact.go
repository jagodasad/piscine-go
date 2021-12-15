package piscine

func Compact(ptr *[]string) int {
	len := 0
	for _, s := range *ptr {
		if s != "" {
			len++
		}
	}
	cou := 0
	arr := make([]string, len)
	for _, s := range *ptr {
		if s != "" {
			arr[cou] = s
			cou++
		}
	}
	*ptr = arr
	return len
}
