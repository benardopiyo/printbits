package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	if len(os.Args) != 2 {
		printStr("00000000 \n")
		return
	}

	input := os.Args[1]
	num := 0

	for _, char := range input {
		if char < '0' || char > '9' {
			printStr("00000000 \n")
			return
		}
		num = num*10 + int(char-'0')
	}

	var binaryStr string
	for i := 7; i >= 0; i-- {
		if num&(1<<i) != 0 {
			binaryStr += "1"
		} else {
			binaryStr += "0"
		}
	}
	printStr(binaryStr + "\n")
}
