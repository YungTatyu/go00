package piscine

import "ft"

func print(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func createThreeDigitComb() string {
	var str string
	for fi := 0; fi <= 7; fi++ {
		for si := fi + 1; si <= 8; si++ {
			for ti := si + 1; ti <= 9; ti++ {
				var newNum string
				newNum += string(fi + '0')
				newNum += string(si + '0')
				newNum += string(ti + '0')
				if newNum != "789" {
					newNum += ", "
				}
				str += newNum
			}
		}
	}
	return str
}

func PrintComb() {
	print(createThreeDigitComb())
}
