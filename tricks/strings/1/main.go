package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) // выводит: true

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) // выводит: false
}
