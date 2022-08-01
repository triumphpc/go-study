package main

import "fmt"

func main() {
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := num[:5:8]
	fmt.Println(s1, cap(s1), len(s1)) //[0 1 2 3 4] 8 5
	//s2 := num[:6:5] // нельзя выделять капасити меньше или больше диапазона

}
