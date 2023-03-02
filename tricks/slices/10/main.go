package main

import "fmt"

func main() {
	str := "Воттак"

	fmt.Println(len(str)) // 12

	s := []rune(str)
	fmt.Println(len(s)) // 6

	ns := append(s[:3], '.', '.', '.')
	fmt.Println(string(ns)) // Вот...
	fmt.Println(len(ns))    // 6

}
