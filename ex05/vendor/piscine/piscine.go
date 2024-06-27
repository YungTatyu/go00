package piscine

import "ft"

func print(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func convert(num int) string {
	var str string
	var tp int
	var op int
	tp = num / 10
	op = num % 10
	str += string(tp + '0')
	str += string(op + '0')
	return str
}

func PrintComb2() {
	for fi := 0; fi < 99; fi++ {
		for si := fi + 1; si < 100; si++ {
			var str string
			str += convert(fi)
			str += " "
			str += convert(si)
			if str != "98 99" {
				str += ", "
			}
			print(str)
		}
	}
}
