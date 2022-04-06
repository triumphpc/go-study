package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "♥"
	fmt.Println(len(data))                    // выводит: 3 (выводит количество байт)
	fmt.Println(utf8.RuneCountInString(data)) // выводит: 1

	// Технически функция RuneCountInString() не возвращает количество символов, потому что один символ может занимать несколько рун.
	data = "é"
	fmt.Println(len(data))                    // выводит: 3
	fmt.Println(utf8.RuneCountInString(data)) // выводит: 2
}
