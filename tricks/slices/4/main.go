package main

import "fmt"

func main() {
	s1 := make([]int, 1)
	fmt.Printf("%v %T %p\n", s1, s1, s1)

	s2 := s1
	s2[0] = 1

	fmt.Printf("%v %T %p\n", s1, s1, s1)
	fmt.Printf("%v %T %p\n", s2, s2, s2)

	// create new address and new pinter
	s1 = append(s1, 44)

	fmt.Printf("%v %T %p\n", s1, s1, s1)
	fmt.Printf("%v %T %p\n", s2, s2, s2) // it's old value

	s2[0] = 5

	fmt.Printf("%v %T %p\n", s1, s1, s1)
	fmt.Printf("%v %T %p\n", s2, s2, s2)

}
