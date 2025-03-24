package piscine

func UltimateDivMod(a *int, b *int) {
	var tempa int = *a
	var tempb int = *b
	*a = tempa / tempb
	*b = tempa % tempb
}
