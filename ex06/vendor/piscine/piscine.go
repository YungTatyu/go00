package piscine

import (
	"ft"
)

func Print(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func generateCombinations(prefix string, start int, n int, combinations *[]string) {
	if len(prefix) == n {
		*combinations = append(*combinations, prefix)
		return
	}

	for i := start; i <= 9; i++ {
		generateCombinations(prefix+string(rune('0'+i)), i+1, n, combinations)
	}
}

func printCombinations(combinations[] string) {
	Print(combinations[0])
	for _, comb := range combinations[1:] {
		Print(", " + comb)
	}
	ft.PrintRune('\n')
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var combinations []string
	generateCombinations("", 0, n, &combinations)
	printCombinations(combinations)
}
