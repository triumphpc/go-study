package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}

	s0 := a[1:]    // [1,2,3,4,5,6] [6, 6]
	s1 := s0[:2:4] // [1,2,x,x] [2, 4]

	fmt.Println(s0)     // [1 2 3 4 5 6]
	fmt.Println(s0[4:]) // [5 6]

	s2 := append(s1, s0[4:]...) // [1,2,4,5,6] - новый массив будет: 0 1 2 5 6 5 6
	s2 = append(s2, 7)          // [1,2,4,5,6,7]

	fmt.Println(a) // [0 1 2 5 6 5 6]

}
