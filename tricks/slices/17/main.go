package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}

	s0 := a[1:]    // [1,2,3,4,5,6] [6, 6]
	s1 := s0[:2:4] // [1,2,x,x] [2, 4]

	fmt.Println(s0)     // [1 2 3 4 5 6]
	fmt.Println(s0[4:]) // [5 6]

	s2 := append(s1, s0[4:]...) // [1,2,5,6]
	fmt.Println(a)              // [0 1 2 5 6 5 6]

	fmt.Println(s2, len(s2), cap(s1)) // [1 2 5 6] 4 4

	s2 = append(s2, 7)                // новый выделяет
	fmt.Println(s2, len(s2), cap(s2)) // [1 2 5 6 7] 5 8

	fmt.Println(a) // [0 1 2 5 6 5 6]

}
