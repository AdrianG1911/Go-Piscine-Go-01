package piscine

func Swap(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
