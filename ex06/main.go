package main

import (
	"ft"
	"piscine"
)

func main() {
	for i := 0; i <= 10; i++ {
		ft.PrintRune(rune(i + '0'))
		piscine.Print("-----arg-----\n")
		piscine.PrintCombN(i)
		piscine.Print("-------------\n")
	}
}
