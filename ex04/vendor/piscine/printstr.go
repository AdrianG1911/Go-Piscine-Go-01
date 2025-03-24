package piscine

import "ft"

func PrintStr(s string) {
	for _, let := range s {
		ft.PrintRune(let)
	}
}
