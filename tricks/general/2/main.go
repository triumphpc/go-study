package main

import "fmt"

func main() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ', '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t')) // выводит true (хорошо)
	fmt.Println(isSpace(' '))  // выводит true (хорошо)
}
