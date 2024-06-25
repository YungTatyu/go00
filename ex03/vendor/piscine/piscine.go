package piscine

import "ft"

func print(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func IsNegative(nb int) {
	var str string
	if nb < 0 {
		str = "T\n"
	} else {
		str = "F\n"
	}
	print(str)
}
