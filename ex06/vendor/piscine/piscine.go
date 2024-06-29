package piscine

import (
	"ft"
)

func Print(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func printRunes(s string) {
	for i, r := range s {
		if i > 0 {
			Print(", ")
		}
		Print(string(r))
	}
}

func generateCombinations(prefix string, start int, n int) {
	if len(prefix) == n {
		Print(prefix)
		return
	}

	for i := start; i <= 9; i++ {
		if len(prefix)+1 == n {
			Print(prefix + string(rune('0'+i)))
			if prefix != "" || i < 9 {
				Print(", ")
			}
		} else {
			generateCombinations(prefix+string(rune('0'+i)), i+1, n)
		}
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	generateCombinations("", 0, n)
	ft.PrintRune('\n')
}
