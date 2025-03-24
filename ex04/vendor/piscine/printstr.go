package piscine

import "ft"

func PrintStr(s string) {
	for _, ch := range s {
		ft.PrintRune(ch)
	}
}
