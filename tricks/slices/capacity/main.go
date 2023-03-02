package main

import "fmt"

func main() {
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := num[:5:8]
	fmt.Printf("%p %s len=%v cap=%v\n", num, num, len(num), cap(num)) // [0 1 2 3 ...] 10 10
	fmt.Printf("%p %s len=%v cap=%v\n", s1, s1, len(s1), cap(s1))     // [0 1 2 3 4 . . .]  5 8

	//s2 := num[::4] нельзя для такого формата укаазывать меньший капасити
	//fmt.Printf("%p %s len=%v cap=%v\n", s2, s2, len(s2), cap(s2))     // [0 1 2 3 4 . . .]  5 8

}
