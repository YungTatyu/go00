package main

import (
	"ex00/vender/ft"
)

func main() {
	for c := 'a'; c <= 'z'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
