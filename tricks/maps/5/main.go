package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 5

	fmt.Println(m)
	change(m) // При передаче мапы в функции - передается всегда указатель на нее
	fmt.Println(m)

	// А что же со слайсом?
	// Слайс - эт ссылочный тип, он передается по значению
	// Но при смене элемента - меняется значение базового массива
	s := make([]int, 1, 3)
	s[0] = 5

	fmt.Println(s)
	changeSlice(s)
	fmt.Println(s)
}

func change(m map[int]int) {
	m[0] = 123
}

func changeSlice(s []int) {
	s[0] = 123
}
