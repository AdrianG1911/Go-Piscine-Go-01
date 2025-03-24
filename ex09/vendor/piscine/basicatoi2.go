package piscine

func isnum(b rune) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func BasicAtoi2(s string) int {
	num := 0
	for _, c := range s {
		if isnum(c) == false {
			return 0
		}
		num = (num * 10) + int(c-'0')
	}
	return num
}
