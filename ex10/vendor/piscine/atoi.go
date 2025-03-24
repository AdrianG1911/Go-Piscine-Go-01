package piscine

func isnum(b rune) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func Atoi(s string) int {
	num := 0
	neg := 1
	if s[0] == '-' {
		neg = -1
		s = s[1:]
	}
	if s[0] == '+' {
		s = s[1:]
	}
	for _, c := range s {
		if isnum(c) == false {
			return 0
		}
		num = (num * 10) + int(c-'0')*neg
	}
	return num
}
