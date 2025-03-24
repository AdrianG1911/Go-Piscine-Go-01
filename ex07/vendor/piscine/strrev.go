package piscine

func strlen(s string) int {
	i := 0
	for range s{
		i++
	}
	return i
}

func StrRev(s string) string {
	srunes := []rune(s)
	len := strlen(s)
	i := 0
	for i < (len - 1) {
		srunes[i], srunes[len-1] = srunes[len-1], srunes[i]
		i++
		len--
	}
	return string(srunes)
}
