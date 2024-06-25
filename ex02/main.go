package main

import (
	"ex02/vender/ft"
)

func main() {
	for c := '0'; c <= '9'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
