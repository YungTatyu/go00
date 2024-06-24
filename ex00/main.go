package main

import (
	"fmt"
)

func main() {
	var message string
	for c := 'a'; c <= 'z'; c++ {
		message += string(c)
	}
	fmt.Println(message)
}
