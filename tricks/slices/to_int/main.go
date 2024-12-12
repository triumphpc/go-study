package main

import "fmt"

// Как конвертировать слайс в инт
func main() {
	// Since 1.17
	//s := []int{1, 2, 3}
	//a := *(*[3]byte)(s)
	//fmt.Println(a)

	// Since 1.20
	//s := []int{1, 2, 3}
	//a := [3]int(s)
	//fmt.Println(a)

	s := make([]byte, 2, 4)
	s0 := (*[0]byte)(s) // s0 != nil
	s2 := (*[2]byte)(s) // &amp;s2[0] == &amp;s[0]
	//s4 := (*[4]byte)(s)
	fmt.Println(s0, s2)

	type A [4]int
	var s5 = (*A)([]int{1, 2, 3, 4})
	fmt.Printf("%T\n", s5)

	slice := []int64{10, 20, 30, 40}
	array := [4]int64(slice)
	fmt.Printf("%T\n", array) // [4]int64
}
