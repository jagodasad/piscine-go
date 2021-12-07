package piscine

func IsRune(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	characters := []rune(s)
	c1 := true

	for i := range characters {
		if IsRune(characters[i]) == true && c1 {
			if 'a' <= characters[i] && characters[i] <= 'z' {
				characters[i] -= 32
			}
			c1 = false
		} else if 'A' <= characters[i] && characters[i] <= 'Z' {
			characters[i] += 32
		} else if IsRune(characters[i]) == false {
			c1 = true
		}
	}
	return string(characters)
}

/*for i, c := range s {
	if 97 <= c && 122 >= c {
		characters[i] = rune(c - 32)
	}
}
return string(characters)
else if {
for d, e := range s {
	if 65 <= e && 90 >= e {
		characters[d] = rune(c + 32)
	}
}
return string(characters)}

	characters := []rune(s)
	l := len(characters)
	for i := 0; i < l; i++{
		characters[i] =
	}*/
