package main

import (
	"ex01/vender/ft"
)

func main() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
