package piscine

func StrRev(s string) string {
	srunes := []rune(s)
	len := len(s)
	i := 0
	for i < (len - 1) {
		srunes[i], srunes[len-1] = srunes[len-1], srunes[i]
		i++
		len--
	}
	return string(srunes)
}
